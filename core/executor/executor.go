package executor

type (
	CommandInput  interface{}
	CommandOutput interface{}

	Command interface {
		Execute(CommandInput) (CommandOutput, error)
	}
)
