package services

// ServiceResponse is the type used to serialize Service data that is returned
// to clients of the sherlock api
type ServiceResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	RepoURL string `json:"repo_url"`
}

// ServiceSerializer is used to serializer a single Service model
// to a used to generate responses from the /services api group
type ServiceSerializer struct {
	Service
}

func (ss *ServiceSerializer) Response() ServiceResponse {
	return ServiceResponse{
		ID:      ss.ID,
		Name:    ss.Name,
		RepoURL: ss.RepoURL,
	}
}

// ServicesSerializer is used to serialize a ServiceModel
// entity to a a format used in http response bodies
type ServicesSerializer struct {
	Services []Service
}

// Response serializes Service models into Service Responses
func (ss *ServicesSerializer) Response() []ServiceResponse {
	services := []ServiceResponse{}
	for _, service := range ss.Services {
		serializer := ServiceSerializer{service}
		services = append(services, serializer.Response())
	}
	return services
}

// Response is a type that allows all data returned from the /service api group to share a consistent structure
type Response struct {
	Services []ServiceResponse `json:"services"`
	Error    string            `json:"error,omitempty"`
}
