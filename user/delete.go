package user

import (
	"github.com/makesitgo/poc-cli/common"
	"github.com/mitchellh/cli"
)

func DeleteCommand() (cli.Command, error) {
	return &deleteCommand{&common.Command{
		Name: "user delete",
	}}, nil
}

type deleteCommand struct {
	*common.Command
}

func (cmd *deleteCommand) Run(args []string) int {
	return 0
}

func (cmd *deleteCommand) Help() string {
	return "user delete help"
}

func (cmd *deleteCommand) Synopsis() string {
	return "user delete synopsis"
}
