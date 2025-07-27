package core

import "gorm.io/gorm"

type (
	Instance struct {
		gorm.Model
		Resource

		Name         string          `json:"name"`
		ProviderType string          `json:"provider_type"` // e.g., "aws", "gcp", "azure"
		Access       *InstanceAccess `json:"instance_access,omitempty"`
	}
	InstanceAccess struct {
		PublicIP  string `json:"public_ip"`
		PrivateIP string `json:"private_ip"`
		SSHKey    string `json:"ssh_key"`
	}
	Resource struct {
		Cores  int `json:"cores,omitempty"`  // Number of CPU cores required
		Memory int `json:"memory,omitempty"` // Memory in MB required
		Disk   int `json:"disk,omitempty"`   // Disk space in MB required
		Iops   int `json:"iops,omitempty"`   // IOPS required
		BwIn   int `json:"bw_in,omitempty"`  // Bandwidth in Mbps for incoming traffic
		BwOut  int `json:"bw_out,omitempty"` // Bandwidth in Mbps for outgoing traffic
	}
	Requirement struct {
		Resource      Resource `json:"resource,omitempty"`       // Resource requirements for the instance
		KernelVersion string   `json:"kernel_version,omitempty"` // Required kernel version
	}
	InstanceToResource struct {
		gorm.Model

		ResourceType string `json:"resource_type"` // e.g., "instance", "database", "service"
		ResourceID   uint   `json:"resource_id"`   // ID of the resource in the database
		InstanceID   uint   `json:"instance_id"`   // ID of the instance in the database
	}
)

func NewInstance() (*Instance, error) {
	return &Instance{}, nil
}

func (instance Instance) GetAvailableInstance(requirement Requirement) (*Instance, error) {
	// Implement logic to find an available instance based on the requirement
	// This could involve checking the provider's API or database records
	return &instance, nil
}
