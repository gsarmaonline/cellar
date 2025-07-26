package deployer

type (
	Deployment interface {
		Deploy() error
	}
)
