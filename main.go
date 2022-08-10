package main

import (
	"log"
	"os"

	"github.com/makesitgo/poc-cli/app"
	"github.com/makesitgo/poc-cli/push"
	"github.com/makesitgo/poc-cli/user"
	"github.com/makesitgo/poc-cli/whoami"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("poc-cli", "1.0.0")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"push":         push.Command,
		"app":          app.DefaultCommand,
		"app init":     app.InitCommand,
		"app create":   app.CreateCommand,
		"app describe": app.DescribeCommand,
		"user create":  user.CreateCommand,
		"user delete":  user.DeleteCommand,
		"user revoke":  user.RevokeCommand,
		"user disable": user.DisableCommand,
		"whoami":       whoami.Command,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
