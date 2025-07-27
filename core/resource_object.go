package core

import "gorm.io/gorm"

type (
	ResourceObject struct {
		gorm.Model

		Host

		Name string   `json:"name"`
		Tags []string `json:"tags,omitempty"`

		Env         string       `json:"env,omitempty"`         // e.g., "production", "staging", "development"
		Requirement *Requirement `json:"requirement,omitempty"` // Resource requirements for the object

		Instances []*Instance `json:"instances,omitempty"` // List of instances associated with this resource object
	}
	Host struct {
		Hostname string `json:"hostname"`
		IP       string `json:"ip,omitempty"`
		Dns      string `json:"dns,omitempty"`  // Optional DNS name
		Port     int    `json:"port,omitempty"` // Optional port number
	}
)
