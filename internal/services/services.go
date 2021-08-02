package services

// Service is the data structure representing an indvidual applicaiton
type Service struct {
	ID      int    `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	RepoURL string `json:"repo_url" db:"repo_url"`
}
