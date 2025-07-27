package core

import "gorm.io/gorm"

type (
	BaseObject struct {
		gorm.Model

		ID string `json:"id"`

		CreatedAt string  `json:"created_at"`
		UpdatedAt string  `json:"updated_at"`
		DeletedAt *string `json:"deleted_at,omitempty"`

		Name string   `json:"name"`
		Tags []string `json:"tags,omitempty"`

		Env string `json:"env,omitempty"` // e.g., "production", "staging", "development"
	}
	Host struct {
		Hostname string `json:"hostname"`
		IP       string `json:"ip,omitempty"`
		Dns      string `json:"dns,omitempty"`  // Optional DNS name
		Port     int    `json:"port,omitempty"` // Optional port number
	}
)
