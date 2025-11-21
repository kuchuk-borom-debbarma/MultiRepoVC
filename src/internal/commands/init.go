package commands

import (
	v1 "MultiRepoVC/src/internal/core/version_control/v1"
	"MultiRepoVC/src/internal/utils/arg"
	"errors"
)

type InitCommand struct{}

func (c *InitCommand) Name() string {
	return "init"
}

func (c *InitCommand) Description() string {
	return "Initializes a new MRVC repository in the current directory."
}

func (c *InitCommand) RequiredArgs() []string {
	return []string{"name", "author"}
}

func (c *InitCommand) OptionalArgs() []string {
	return []string{}
}

func (c *InitCommand) Execute(args []string) error {
	parsed := arg.ParseArgs(args)

	name := parsed["name"]
	author := parsed["author"]

	if name == "" {
		return errors.New("missing required argument: --name")
	}
	if author == "" {
		return errors.New("missing required argument: --author")
	}

	vc := v1.New()
	return vc.Init(name, author)
}

func init() {
	Global.Register(&InitCommand{})
}
