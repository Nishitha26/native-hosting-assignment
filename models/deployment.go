package models

import "time"

type Deployment struct {
	SiteName   string    `json:"site_name"`
	DeployedAt time.Time `json:"deployed_at"`
	FileCount  int       `json:"file_count"`
}
