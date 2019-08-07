import matplotlib.pyplot as plt
from benckmark.controller import Controller
from benchmark.satellite import MockSatelliteGroup as SatelliteGroup
import numpy as np
import argparse
from os import path
from threading import Timer
import logging
import time

TRIALS = 2
TRIAL_LENGTH = 90
DISCONNECT_TIME = 30
RECONNECT_TIME = 60

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Produce graphs of what happens when satellites disconnect / reconnect.')
    parser.add_argument('client', help='Name of the client to use in these tests.')
    args = parser.parse_args()

    with Controller(args.client) as controller:
        # two stacked plots in one figure
        fig, ax = plt.subplots(1, 4, sharex='col', sharey='row', figsize=(20, 8), dpi=100)
        fig.suptitle(f'{controller.client_name.title()} Satellite Disconnect')

        ax[0].set_title("Untraced")
        ax[0].set(xlabel="Time (s)", ylabel="Memory (MB)")

        # setup satellite disconnect column
        ax[1].set_title("Traced")
        ax[1].set_xlabel("Time (s)")

        # setup nominal column
        ax[2].set_title("Satellite Disconnect")
        ax[2].set_xlabel("Time (s)")
        ax[2].axvline(x=DISCONNECT_TIME, alpha=.2, color='red')

        # setup restart column
        ax[3].set_title("Satellite Reconnect")
        ax[3].set_xlabel("Time (s)")
        ax[3].axvline(x=DISCONNECT_TIME, alpha=.2, color='red')
        ax[3].axvline(x=RECONNECT_TIME, alpha=.2, color='green')

        for i in range(TRIALS):
            for index, trace, action in [
                    (0, False, None),
                    (1, True, None),
                    (2, True, 'disconnect'),
                    (3, True, 'reconnect')]:

                # Don't initialize using a with statement because we are going
                # to shut this down manually.
                satellites = SatelliteGroup('typical')
                logging.info("trial {} traced {} type {}".format(index, trace, action))

                def satellite_action():
                    if action == 'disconnect':
                        logging.info("shutting down")
                        satellites.shutdown()
                    if action == 'reconnect':
                        logging.info("shutting down")
                        satellites.shutdown()
                        time.sleep(RECONNECT_TIME - DISCONNECT_TIME)
                        satellites.start('typical')
                        logging.info("reconnected")

                # satellites shutdown in the middle of the test
                shutdown_timer = Timer(DISCONNECT_TIME, satellite_action)
                shutdown_timer.start()

                result = controller.benchmark(
                    trace=trace,
                    spans_per_second=50,
                    runtime=TRIAL_LENGTH,
                )

                print(result)

                logging.info("benchmark completed")

                if action != 'disconnect':
                    satellites.shutdown()

                sample_time = list(range(1, len(result.cpu_list) + 1))
                ax[index].plot(sample_time, [m * 2**-20 for m in result.memory_list])

        fig.savefig(f'graphs/{controller.client_name}_satellite_disconnect.png')
