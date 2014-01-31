go-cf-autoscaler
================

Demonstration of autoscaling workers in a producer-consumer scenario on Cloud Foundry....in Go!

[![baby-gopher](https://raw2.github.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

This project is major WIP! I am attempting to port my Java/Spring [cf-autoscaler](https://github.com/cloudfoundry-samples/cf-autoscaler) to Go as a learning exercise.

Stay tuned for more goodness.

Current State
-------------

* Hardcoded to connect to a local RabbitMQ instance
* Get it: `go get github.com/mstine/go-cf-autoscaler`
* Build it: `go install`
* Run the producer: `go-cf-autoscaler p`
* Run the worker: `go-cf-autoscaler w`
* Get help: `go-cf-autoscaler help`
