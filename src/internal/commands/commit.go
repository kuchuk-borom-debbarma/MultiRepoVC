package commands

import (
	v1 "MultiRepoVC/src/internal/core/version_control/v1"
	"MultiRepoVC/src/internal/utils/arg"
	"errors"
	"fmt"
)

type CommitCommand struct{}

func (c *CommitCommand) Name() string {
	return "commit"
}

func (c *CommitCommand) Description() string {
	return "Creates a new commit with the specified message and files."
}

func (c *CommitCommand) RequiredArgs() []string {
	return []string{"message"}
}

func (c *CommitCommand) OptionalArgs() []string {
	return []string{"author"}
}

// -------------------------------------------------------
// 🎯 Execution
// -------------------------------------------------------

func (c *CommitCommand) Execute(args []string) error {
	p := arg.ParseArgs(args)

	// Required: --message
	message := p["message"]
	if message == "" {
		return errors.New("missing required argument: --message")
	}

	// Optional: --author
	author := p["author"]
	if author == "" {
		author = "unknown"
	}

	// Positional files (0, 1, 2…)
	var files []string
	for i := 0; ; i++ {
		key := fmt.Sprintf("%d", i)
		val, ok := p[key]
		if !ok {
			break
		}
		files = append(files, val)
	}

	if len(files) == 0 {
		return errors.New("no files provided for commit")
	}

	vc := v1.New()
	return vc.Commit(message, author, files)
}

// -------------------------------------------------------
// 🎯 Register this command
// -------------------------------------------------------
func init() {
	Global.Register(&CommitCommand{})
}
