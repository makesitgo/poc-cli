package app

import (
	"github.com/makesitgo/poc-cli/common"
	"github.com/mitchellh/cli"
)

func CreateCommand() (cli.Command, error) {
	return &createCommand{&common.Command{
		Name: "app create",
	}}, nil
}

type createCommand struct {
	*common.Command
}

func (cmd *createCommand) Run(args []string) int {
	return 0
}

func (cmd *createCommand) Help() string {
	return "app create help"
}

func (cmd *createCommand) Synopsis() string {
	return "app create synopsis"
}
