go-cf-autoscaler
================

Demonstration of autoscaling workers in a producer-consumer scenario on Cloud Foundry....in Go!

[![baby-gopher](https://raw2.github.com/drnic/babygopher-site/gh-pages/images/babygopher-badge.png)](http://www.babygopher.org)

This project is major WIP! I am attempting to port my Java/Spring [cf-autoscaler](https://github.com/cloudfoundry-samples/cf-autoscaler) to Go as a learning exercise.

Stay tuned for more goodness.

Current State
-------------

* Looks for an AMQP URI in environment variable `VCAP_SERVICES`:

```
{
    "cloudamqp-n/a": [
        {
            "credentials": {
                "uri": "amqp://user:********@green-skunk.rmq.cloudamqp.com/xweijkex"
            }, 
            "label": "cloudamqp-n/a", 
            "name": "cloudamqp-autoscale", 
            "plan": "bunny", 
            "tags": [
                "amqp", 
                "rabbitmq"
            ]
        }
    ]
}
```

* Get it: `go get github.com/mstine/go-cf-autoscaler`
* Build it: `go install`
* Run the producer: `go-cf-autoscaler p`
* Run the worker: `go-cf-autoscaler w`
* Get help: `go-cf-autoscaler help`
