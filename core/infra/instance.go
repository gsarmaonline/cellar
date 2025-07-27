package infra

type (
	Instance struct {
		Name         string          `json:"name"`
		ProviderType string          `json:"provider_type"` // e.g., "aws", "gcp", "azure"
		Access       *InstanceAccess `json:"instance_access,omitempty"`
	}

	InstanceAccess struct {
		PublicIP  string `json:"public_ip"`
		PrivateIP string `json:"private_ip"`
		SSHKey    string `json:"ssh_key"`
	}
)
