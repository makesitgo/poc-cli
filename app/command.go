package app

import (
	"github.com/mitchellh/cli"
)

func DefaultCommand() (cli.Command, error) {
	return &defaultCommand{newInitCommand()}, nil
}

type defaultCommand struct {
	*initCommand
}

func (cmd *defaultCommand) Synopsis() string {
	return "perform app-related actions, default sub-command: init"
}

func (cmd *defaultCommand) Help() string {
	return `The app command allows you to perform various app-related actions.

To do so, run this command with one of the following sub-commands:`
}
