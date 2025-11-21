package commands

// Command is the main interface for all MRVC CLI commands.
// BaseCommand.Run() will call RequiredArgs() and OptionalArgs()
// and validate before calling ExecuteCommand().
type Command interface {
	Name() string
	Description() string
	RequiredArgs() []string
	OptionalArgs() []string

	// ExecuteCommand Called AFTER args are validated & parsed
	ExecuteCommand(parsed map[string]string) error
}
