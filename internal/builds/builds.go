package builds

import "time"

type Build struct {
	ID            int       `json:"id,omitempty"`
	VersionString string    `json:"version_string" binding:"required"`
	CommitSha     string    `json:"commit_sha" binding:"required"`
	BuildURL      string    `json:"build_url,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	ServiceID     int       `json:"service_id" binding:"required"`
}

type BuildModel interface {
	ListAll() ([]Build, error)
}
