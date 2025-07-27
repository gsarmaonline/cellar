package interfaces

type (
	CommandInput  map[string]interface{}
	CommandOutput map[string]interface{}

	Command interface {
		Execute(CommandInput) (CommandOutput, error)
	}
)
