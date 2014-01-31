package cf

import (
	"encoding/json"
	"os"
	"fmt"
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

func SingleAmqpUri() string {
	servicesJson := []byte(os.Getenv("VCAP_SERVICES"))
	var services interface{}
	err := json.Unmarshal(servicesJson, &services)
	util.FailOnError(err, "Failed to unmarshal VCAP_SERVICES")

	servicesMap := services.(map[string]interface{})

	cloudamqpArray := servicesMap["cloudamqp-n/a"].([]interface{})
	cloudamqpService := cloudamqpArray[0].(map[string]interface{})
	credentials := cloudamqpService["credentials"].(map[string]interface{})
	uri := credentials["uri"].(string)

	fmt.Printf("URI: %v", uri)
	return uri
}
