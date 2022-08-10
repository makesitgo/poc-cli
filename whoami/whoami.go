package whoami

import (
	"github.com/makesitgo/poc-cli/common"

	"github.com/mitchellh/cli"
)

func Command() (cli.Command, error) {
	return &command{&common.Command{
		Name:        "whoami",
		Description: "view details related to your current session",
		Requester: func(flagSet common.Flags) common.Requester {
			var flags flags
			flagSet.StringVar(&flags.dummyFlag, "dummy-flag", "", "")

			return requester{flags, prompts{}}
		},
	}}, nil
}

type command struct {
	*common.Command
}

type requester struct {
	flags
	prompts
}

func (rr requester) Request() (interface{}, error) {
	return nil, nil
}

type prompts struct {
}

func (cmd *command) Run(args []string) int {
	// if cmd.flags == nil {
	// 	cmd.flags = newFlags(cmd.Name)
	// }
	//
	// if err := cmd.flags.Parse(args); err != nil {
	// 	fmt.Println("failed to parse flags", err)
	// 	return 1
	// }

	/*
		if c.user != nil {
			return c.user, nil
		}

		u, err := c.storage.ReadUserConfig()
		if err != nil {
			return nil, err
		}

		c.user = u

		return u, nil
	*/

	// if cmd.user == nil {
	//
	// }
	return 0
}

type flags struct {
	common.Flags
	dummyFlag string
}

func newFlags(name string) flags {
	var flags flags

	// flagSet := common.NewFlags(name)
	// flagSet.StringVar(&flags.dummyFlag, "dummy-flag", "", "")

	// flags.Flags = flagSet
	return flags
}
