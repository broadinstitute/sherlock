package environments

import "github.com/broadinstitute/sherlock/internal/models"

// EnvironmentResponse is the type that environment
// entities are serialized to and used in responses to clients
type EnvironmentResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type EnvironmentSerializer struct {
	Environment models.Environment
}

func (es *EnvironmentSerializer) Response() EnvironmentResponse {
	return EnvironmentResponse{
		Name: es.Environment.Name,
		ID:   es.Environment.ID,
	}
}

// EnvironmentsSerializer is used to convert a list of environment model types
// to a Respopnse type used to send environment info to clients
type EnvironmentsSerializer struct {
	Environments []models.Environment
}

// Response is a function that Serializers a slice of Environment models
// to responses suitable for sending to clients
func (es *EnvironmentsSerializer) Response() []EnvironmentResponse {
	environments := []EnvironmentResponse{}
	for _, environment := range es.Environments {
		serializer := EnvironmentSerializer{environment}
		environments = append(environments, serializer.Response())
	}
	return environments
}
