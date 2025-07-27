package core

import "github.com/gsarmaonline/cellar/core/infra"

const (
	ExecutorTypeDocker = "docker"
	ExecutorTypeGolang = "golang"
)

type (
	ExecutorTypeT string

	Service struct {
		BaseObject

		Parallelism  int               `json:"parallelism,omitempty"` // Number of parallel instances to run
		ExecutorType ExecutorTypeT     `json:"executor_type"`         // e.g., "docker", "kubernetes", "local"
		Instances    []*infra.Instance `json:"instance,omitempty"`
	}
)
