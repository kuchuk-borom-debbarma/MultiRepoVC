package commands

import (
	v1 "MultiRepoVC/src/internal/core/version_control/v1"
	"errors"
	"fmt"
)

type CommitCommand struct {
	BaseCommand
}

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

func (c *CommitCommand) ExecuteCommand(p map[string]string) error {
	message := p["message"]

	author := p["author"]
	if author == "" {
		author = "unknown"
	}

	// Extract positional file args (0,1,2,...)
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

func init() {
	Global.Register(&CommitCommand{})
}
