"use strict";

const opentracing = require("opentracing");
const lightstep = require("lightstep-tracer");
const http = require("http");

const satellite_host = "localhost";
const satellite_port = 8360;

const controller_port = 8023;
const controller_host = "localhost";

const millis_per_nano = 1000000;
const prime_work = 982451653;

const noop_tracer = new opentracing.Tracer(null);
const test_tracer = new lightstep.Tracer({
  access_token: "ignored",
  collector_port: satellite_port,
  collector_host: satellite_host,
  collector_encryption: "none",
  component_name: "javascript/test",
  report_timeout_millis: 200 // .2s
});

function exec_control(args, tracer) {
  var begin = process.hrtime();
  var sleep_debt = 0;
  var sleep_at;
  var sleep_nanos = 0;
  var p = prime_work;
  var body_func = function(repeat) {
    if (sleep_debt > 0) {
      var diff = process.hrtime(sleep_at);
      var nanos = diff[0] * 1e9 + diff[1];
      sleep_nanos += nanos;
      sleep_debt -= nanos;
    }
    for (var r = repeat; r > 0; r--) {
      var span = tracer.startSpan("span/test");
      for (var i = 0; i < args.work; i++) {
        p *= p;
        p %= 2147483647;
      }
      span.finish();
      sleep_debt += args.sleep;
      if (sleep_debt > args.sleep_interval) {
        sleep_at = process.hrtime();
        return setTimeout(body_func, sleep_debt / millis_per_nano, r - 1);
      }
    }

    // done with work at this point

    var endTime = process.hrtime();
    var elapsed = endTime[0] - begin[0] + (endTime[1] - begin[1]) * 1e-9;

    if (args.trace && !args.no_flush) {
      // call next_control when the flush is complete
      return tracer.flush();
    }
    process.exit();
  };

  body_func(args.repeat);
}

const args = require('minimist')(process.argv.slice(2));
