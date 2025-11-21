package commands

import (
	"MultiRepoVC/src/internal/utils/arg"
	"fmt"
)

// BaseCommand acts like an abstract parent class (Java style).
// It defines the call order and handles boilerplate logic.
type BaseCommand struct{}

// Run executes a command using the Template Method Pattern:
// 1. Parse arguments
// 2. Validate required arguments
// 3. Call ExecuteCommand() implemented by the actual command
func (b *BaseCommand) Run(cmd Command, args []string) error {
	parsed := arg.ParseArgs(args)

	// Validate required arguments
	for _, req := range cmd.RequiredArgs() {
		if parsed[req] == "" {
			return fmt.Errorf("missing required argument: --%s", req)
		}
	}

	return cmd.ExecuteCommand(parsed)
}
