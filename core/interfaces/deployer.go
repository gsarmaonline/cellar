package interfaces

type (
	Deployment interface {
		Deploy() error
		GetConfig() (DeploymentConfig, error)
	}
	DeploymentConfig struct {
		Meta map[string]interface{} `json:"meta,omitempty"`
		Dns  string                 `json:"dns,omitempty"`  // DNS name for the deployment
		Port string                 `json:"port,omitempty"` // Port for the deployment
	}
)
