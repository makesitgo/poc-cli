package push

import (
	"github.com/makesitgo/poc-cli/common"
	"github.com/mitchellh/cli"
)

func Command() (cli.Command, error) {
	return &command{&common.Command{
		Name: "push",
	}}, nil
}

type command struct {
	*common.Command
}

func (cmd *command) Run(args []string) int {
	return 0
}

func (cmd *command) Help() string {
	return "push help"
}

func (cmd *command) Synopsis() string {
	return "push an app's changes to Realm"
}
