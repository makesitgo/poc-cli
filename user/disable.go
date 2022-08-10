package user

import (
	"github.com/makesitgo/poc-cli/common"
	"github.com/mitchellh/cli"
)

func DisableCommand() (cli.Command, error) {
	return &disableCommand{&common.Command{
		Name: "user disable",
	}}, nil
}

type disableCommand struct {
	*common.Command
}

func (cmd *disableCommand) Run(args []string) int {
	return 0
}

func (cmd *disableCommand) Help() string {
	return "user disable help"
}

func (cmd *disableCommand) Synopsis() string {
	return "user disable synopsis"
}
