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
	}
)

func (s *Service) Deploy() error {
	// Implement deployment logic here
	// This could involve creating instances, setting up networking, etc.
	return nil
}
