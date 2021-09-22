package environments

// EnvironmentResponse is the type that environment
// entities are serialized to and used in responses to clients
type EnvironmentResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type environmentSerializer struct {
	environment Environment
}

func (es *environmentSerializer) Response() EnvironmentResponse {
	return EnvironmentResponse{
		Name: es.environment.Name,
		ID:   es.environment.ID,
	}
}

// EnvironmentsSerializer is used to convert a list of environment model types
// to a Respopnse type used to send environment info to clients
type EnvironmentsSerializer struct {
	Environments []Environment
}

// Response is a function that Serializers a slice of Environment models
// to responses suitable for sending to clients
func (es *EnvironmentsSerializer) Response() []EnvironmentResponse {
	environments := []EnvironmentResponse{}
	for _, environment := range es.Environments {
		serializer := environmentSerializer{environment}
		environments = append(environments, serializer.Response())
	}
	return environments
}
