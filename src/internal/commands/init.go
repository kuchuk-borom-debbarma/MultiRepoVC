package commands

import (
	v1 "MultiRepoVC/src/internal/core/version_control/v1"
	"MultiRepoVC/src/internal/utils/arg"
)

type InitCommand struct{}

func (c *InitCommand) Name() string {
	return "init"
}

func (c *InitCommand) Execute(args []string) error {
	parsedArgs := arg.ParseArgs(args)
	vc := v1.New()
	return vc.Init(parsedArgs["name"], parsedArgs["author"])
}

// auto-register this command
func init() {
	Register(&InitCommand{})
}
