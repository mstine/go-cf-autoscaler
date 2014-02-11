package cf

import (
	"encoding/json"
	"os"
	"github.com/mstine/go-cf-autoscaler/util"
	"github.com/mitchellh/mapstructure"
)

type Credential struct {
	Uri string
}

type Service struct {
	Credentials Credential
}

func SingleUri(serviceType string) string {	
	services := BoundServices()	
	return services[serviceType][0].Credentials.Uri
}

func BoundServices() map[string][]Service {
	servicesJson := []byte(os.Getenv("VCAP_SERVICES"))
	var services map[string]interface{}
	err := json.Unmarshal(servicesJson, &services)
	util.FailOnError(err, "Failed to unmarshal VCAP_SERVICES")

	boundServices := make(map[string][]Service)
	for k, v := range services {
		var serviceInstances []Service
		err := mapstructure.Decode(v, &serviceInstances)
		util.FailOnError(err, "Failed to unmarshal service: " + k)
		boundServices[k] = serviceInstances
	}	
	return boundServices
}
