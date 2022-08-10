package user

import (
	"github.com/makesitgo/poc-cli/common"
	"github.com/mitchellh/cli"
)

func RevokeCommand() (cli.Command, error) {
	return &revokeCommand{&common.Command{
		Name: "user revoke",
	}}, nil
}

type revokeCommand struct {
	*common.Command
}

func (cmd *revokeCommand) Run(args []string) int {
	return 0
}

func (cmd *revokeCommand) Help() string {
	return "user revoke help"
}

func (cmd *revokeCommand) Synopsis() string {
	return "user revoke synopsis"
}
