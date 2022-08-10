package app

import (
	"github.com/makesitgo/poc-cli/common"
	"github.com/mitchellh/cli"
)

func InitCommand() (cli.Command, error) {
	return newInitCommand(), nil
}

func newInitCommand() *initCommand {
	return &initCommand{&common.Command{
		Name: "app init",
	}}
}

type initCommand struct {
	*common.Command
}

func (cmd *initCommand) Run(args []string) int {
	return 0
}

func (cmd *initCommand) Help() string {
	return "app init help"
}

func (cmd *initCommand) Synopsis() string {
	return "app init synopsis"
}
