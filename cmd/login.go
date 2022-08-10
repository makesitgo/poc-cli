package cmd

import (
	"flag"

	"github.com/makesitgo/poc-cli/internal/common"

	"github.com/mitchellh/cli"
)

func Login() (cli.Command, error) {
	return &common.Command{
		Meta: common.CommandMeta{
			Name:        "login",
			Description: "Authenticate with an Atlas programmatic API Key",
		},
	}, nil
}

type loginFlags struct {
	common.CommandFlags
	apiKey        string
	privateAPIKey string
}

func (flags loginFlags) NewFlagSet(name string) *flag.FlagSet {
	set := flags.CommandFlags.NewFlagSet(name)
	set.StringVar(&flags.apiKey, "api-key", "", "")
	set.StringVar(&flags.privateAPIKey, "private-api-key", "", "")
	return set
}

type loginInputs struct {
	apiKey        string
	privateAPIKey string
}

type loginCommand struct {
	*common.Command
	inputs loginInputs
}

func (cmd *loginCommand) CollectInputs() error {
	cmd.inputs = loginInputs{}
	return nil
}

func (cmd *loginCommand) Run(ctx *common.Context) error {
	return nil
}
