package commands

type Command interface {
	Name() string
	Execute(args []string) error

	// RequiredArgs New: describe arguments
	RequiredArgs() []string // e.g. []{"name", "author"}
	OptionalArgs() []string // e.g. []{"verbose", "force"}

	// Description Short description for help output
	Description() string
}
