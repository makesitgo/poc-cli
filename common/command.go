package common

import (
	"flag"
	"fmt"
)

type Command struct {
	Name        string
	Description string
	Requester   func(flags Flags) Requester
	Worker      func(req interface{}) (interface{}, error)
}

func (cmd *Command) Synopsis() string {
	return cmd.Description
}

func (cmd *Command) Help() string {
	return fmt.Sprintf("The %s command allows you to %s.", cmd.Name, cmd.Description)
}

func (cmd *Command) NewFlags() Flags {
	var flags Flags
	flagSet := flag.NewFlagSet(cmd.Name, flag.ContinueOnError)
	flagSet.StringVar(&flags.configPath, "config-path", "", "")
	flags.FlagSet = flagSet
	return flags
}

func (cmd *Command) Run(args []string) int {
	requester := cmd.Requester(cmd.NewFlags())
	if err := requester.Parse(args); err != nil {
		return 1
	}

	req, reqErr := requester.Request()
	if reqErr != nil {
		return 1
	}

	res, resErr := cmd.Worker(req)
	if resErr != nil {
		return 1
	}

	fmt.Println(res)
	return 0
}

type Requester interface {
	Parse(args []string) error
	Request() (interface{}, error)
}

type Flags struct {
	*flag.FlagSet

	configPath string
}
