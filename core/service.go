package core

const (
	ExecutorTypeDocker = "docker"
	ExecutorTypeGolang = "golang"
)

type (
	ExecutorTypeT string

	Service struct {
		ResourceObject

		Parallelism  int           `json:"parallelism,omitempty"` // Number of parallel instances to run
		ExecutorType ExecutorTypeT `json:"executor_type"`         // e.g., "docker", "kubernetes", "local"
		Instances    []*Instance   `json:"instance,omitempty"`

		Requirement Requirement `json:"instance_requirement,omitempty"` // Requirements for the service instance
	}
)

func (s *Service) Deploy() error {
	// Implement deployment logic here
	// This could involve creating instances, setting up networking, etc.
	return nil
}
