package cf

import (
	"encoding/json"
	"os"
	"github.com/mstine/go-cf-autoscaler/util"
)

/*
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
*/

type Credential struct {
	Uri string `json:"uri"`
}

type Service struct {
	Credentials Credential `json:"credentials"`
}

type AmqpServices struct {
	Services []Service `json:"cloudamqp-n/a"`
}

func SingleAmqpUri() string {
	servicesJson := []byte(os.Getenv("VCAP_SERVICES"))
	var services AmqpServices
	err := json.Unmarshal(servicesJson, &services)
	util.FailOnError(err, "Failed to unmarshal VCAP_SERVICES")
	
	return services.Services[0].Credentials.Uri
}
