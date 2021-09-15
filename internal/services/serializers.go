package services

// ServiceResponse is the type used to serialize Service data that is returned
// to clients of the sherlock api
type ServiceResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	RepoURL string `json:"repo_url"`
}

type ServiceSerializer struct {
	Service
}

func (ss *ServiceSerializer) response() ServiceResponse {
	return ServiceResponse{
		ID:      ss.ID,
		Name:    ss.Name,
		RepoURL: ss.RepoURL,
	}
}

type ServicesSerializer struct {
	Services []Service
}

// Response serializes Service models into Service Responses
func (ss *ServicesSerializer) Response() []ServiceResponse {
	services := []ServiceResponse{}
	for _, service := range ss.Services {
		serializer := ServiceSerializer{service}
		services = append(services, serializer.response())
	}
	return services
}

// Response is a type that allows all data returned from the /service api group to share a consistent structure
type Response struct {
	Services []ServiceResponse `json:"services"`
	Error    string            `json:"error,omitempty"`
}
