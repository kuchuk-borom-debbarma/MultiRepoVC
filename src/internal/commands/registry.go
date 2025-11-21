package commands

import (
	"fmt"
	"strings"
)

var Global = NewRegistry()

type Registry struct {
	commands map[string]Command
}

func NewRegistry() *Registry {
	return &Registry{
		commands: make(map[string]Command),
	}
}

func (r *Registry) Register(cmd Command) {
	r.commands[cmd.Name()] = cmd
}

func (r *Registry) Get(name string) (Command, error) {
	cmd, ok := r.commands[name]
	if !ok {
		return nil, fmt.Errorf("unknown command: %s", name)
	}
	return cmd, nil
}

// List all commands
func (r *Registry) List() {
	fmt.Println("\nAvailable commands:\n")

	for name, cmd := range r.commands {
		fmt.Printf("  %s  -  %s\n", name, cmd.Description())

		// Required args
		req := cmd.RequiredArgs()
		if len(req) > 0 {
			fmt.Printf("      Required: --%s\n", strings.Join(req, " --"))
		}

		// Optional args
		opt := cmd.OptionalArgs()
		if len(opt) > 0 {
			fmt.Printf("      Optional: --%s\n", strings.Join(opt, " --"))
		}

		fmt.Println()
	}
}
