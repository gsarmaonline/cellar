package executors

import "github.com/gsarmaonline/cellar/core/interfaces"

type (
	DockerExecutor struct {
		Image         string            `json:"image"`              // Docker image to use
		ContainerID   string            `json:"container_id"`       // ID of the running container
		ContainerIP   string            `json:"container_ip"`       // IP address of the container
		ContainerPort int               `json:"container_port"`     // Port exposed by the container
		EnvVars       map[string]string `json:"env_vars,omitempty"` // Environment variables for the container
		Volumes       []Volume          `json:"volumes,omitempty"`  // Volumes to mount in the container
	}
	Volume struct {
		Name string `json:"name"` // Name of the volume
		Path string `json:"path"` // Path in the container where the volume is mounted
	}
)

func NewDockerExecutor(image string, envVars map[string]string) *DockerExecutor {
	return &DockerExecutor{
		Image:         image,
		EnvVars:       envVars,
		ContainerID:   "",
		ContainerIP:   "",
		ContainerPort: 0,
	}
}

func (de *DockerExecutor) StartContainer() error {
	// Logic to start a Docker container using the specified image and environment variables
	// This would typically involve using a Docker client library to interact with the Docker API
	return nil
}

func (de *DockerExecutor) StopContainer() error {
	// Logic to stop the Docker container
	// This would typically involve using a Docker client library to stop the container by its ID
	return nil
}

func (de *DockerExecutor) ExecuteCommand(commandInput interfaces.CommandInput) (interfaces.CommandOutput, error) {
	// Logic to execute a command inside the Docker container
	// This would typically involve using a Docker client library to run a command in the container
	return nil, nil
}
