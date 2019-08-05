from controller import Controller, Command, Result
from satellite import SatelliteGroup
import matplotlib.pyplot as plt
import numpy as np
import pytest

@pytest.fixture(scope='module')
def satellites():
    with SatelliteGroup('typical') as satellites:
        yield satellites


## VANILLA TESTS ##

class TestVanilla:
    @pytest.fixture(scope='class')
    def client(self):
        with Controller('python') as controller:
            yield controller

    def test_memory(self, client, satellites):
        """ Tracers should not have memory leaks """

        result_100s = client.benchmark(
            trace=True,
            spans_per_second=500,
            runtime=100,
            satellites=satellites)

        result_5s = client.benchmark(
            trace=True,
            spans_per_second=500,
            runtime=5,
            satellites=satellites)

        assert(result_100s.memory > result_5s.memory)
        # 100s memory < 1.5x 5s memory
        assert(result_100s.memory / result_5s.memory < 1.5)

    def test_dropped_spans(self, client, satellites):
        """ No tracer should drop spans if we're only sending 300 / s. """

        sps_100 = client.benchmark(
            trace=True,
            spans_per_second=100,
            runtime=10,
            satellites=satellites)

        sps_300 = client.benchmark(
            trace=True,
            spans_per_second=300,
            runtime=10,
            satellites=satellites)

        assert(sps_100.dropped_spans == 0)
        assert(sps_300.dropped_spans == 0)

    def test_cpu(self, client, satellites):
        """ Traced ciode shouldn't consume significatly more CPU than untraced
        code """

        TRIALS = 5
        cpu_traced = []
        cpu_untraced = []

        for i in range(TRIALS):
            result_untraced = client.benchmark(
                trace=False,
                spans_per_second=False,
                runtime=10)
            cpu_traced.append(result_traced.cpu_usage * 100)

            result_traced = client.benchmark(
                trace=True,
                spans_per_second=500,
                runtime=10,
                satellites=satellites)
            cpu_untraced.append(result_traced.cpu_usage * 100)

        assert(abs(np.mean(cpu_traced) - np.mean(cpu_untraced)) < 10)
