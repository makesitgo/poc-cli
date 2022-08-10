package app

import (
	"github.com/makesitgo/poc-cli/common"
	"github.com/mitchellh/cli"
)

func DescribeCommand() (cli.Command, error) {
	return &describeCommand{&common.Command{
		Name: "app describe",
	}}, nil
}

type describeCommand struct {
	*common.Command
}

func (cmd *describeCommand) Run(args []string) int {
	return 0
}

func (cmd *describeCommand) Help() string {
	return "app describe help"
}

func (cmd *describeCommand) Synopsis() string {
	return "describe the application structure"
}
