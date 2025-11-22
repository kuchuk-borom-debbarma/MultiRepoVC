package commands

import (
	v1 "MultiRepoVC/src/internal/core/version_control/v1"
	"fmt"
)

type StatusCommand struct {
	BaseCommand
}

func (c *StatusCommand) Name() string { return "status" }

func (c *StatusCommand) Description() string {
	return "Shows the working directory status compared to HEAD."
}

func (c *StatusCommand) RequiredArgs() []string { return []string{} }
func (c *StatusCommand) OptionalArgs() []string { return []string{} }

func (c *StatusCommand) ExecuteCommand(p map[string][]string) error {
	vc := v1.New()
	out, err := vc.Status()
	if err != nil {
		return err
	}

	fmt.Println(out)
	return nil
}

func init() {
	Global.Register(&StatusCommand{})
}
