package user

import (
	"flag"
	"fmt"
	"strings"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/mitchellh/cli"
)

func CreateCommand() (cli.Command, error) {
	return &createCommand{
		flags: NewFlags("user create"),
	}, nil
}

type createCommand struct {
	flags Flags
	mode  string
}

const (
	CreateModeUserPassword = "userpass"
	CreateModeAPIKey       = "api_key"
)

func (cmd *createCommand) Run(args []string) int {
	if err := cmd.flags.Parse(args); err != nil {
		fmt.Println("failed to parse flags", err)
		return 1
	}

	fmt.Println("npmnpm	")

	var colors []string
	if err := survey.AskOne(&survey.MultiSelect{
		Message: "What are you favorite colors?",
		Options: []string{"red", "green", "blue", "orange", "purple", "yellow"},
	}, &colors, survey.WithIcons(func(icons *survey.IconSet) {
		icons.SelectFocus.Format = "cyan+b"
		icons.MarkedOption.Format = "green"
	})); err != nil {
		fmt.Println("failed to ask favorite color question", err)
		return 1
	}
	fmt.Println("favorite colors are", strings.Join(colors, ", "))

	if cmd.flags.username == "" && cmd.flags.publicAPIKey == "" {
		if err := survey.AskOne(&survey.Select{
			Message: "Which auth provider would you like to create a user for?",
			Options: []string{CreateModeUserPassword, CreateModeAPIKey},
			Default: CreateModeUserPassword,
		}, &cmd.mode); err != nil {
			fmt.Println("failed to ask mode question", err)
			return 1
		}
	} else if cmd.flags.publicAPIKey != "" {
		if cmd.flags.username != "" {
			return 1
		}
		cmd.mode = CreateModeAPIKey
	}

	switch cmd.mode {
	case CreateModeUserPassword:
		var questions []*survey.Question
		if cmd.flags.username == "" {
			questions = append(questions, &survey.Question{
				Name:   "username",
				Prompt: &survey.Input{Message: "username:"},
			})
		}
		if cmd.flags.password == "" {
			questions = append(questions, &survey.Question{
				Name:   "password",
				Prompt: &survey.Input{Message: "password:"},
			})
		}

		if len(questions) != 0 {
			answers := struct {
				Username string
				Password string
			}{}

			if err := survey.Ask(questions, &answers); err != nil {
				fmt.Println("failed to ask userpass questions", err)
				return 1
			}

			if answers.Username != "" {
				cmd.flags.username = answers.Username
			}
			if answers.Password != "" {
				cmd.flags.password = answers.Password
			}
		}
	case CreateModeAPIKey:
		var questions []*survey.Question
		if cmd.flags.publicAPIKey == "" {
			questions = append(questions, &survey.Question{
				Name:   "publicAPIKey",
				Prompt: &survey.Input{Message: "public api key:"},
			})
		}
		if cmd.flags.privateAPIKey == "" {
			questions = append(questions, &survey.Question{
				Name:   "privateAPIKey",
				Prompt: &survey.Input{Message: "private api key:"},
			})
		}

		if len(questions) != 0 {
			answers := struct {
				PublicAPIKey  string
				PrivateAPIKey string
			}{}

			if err := survey.Ask(questions, &answers); err != nil {
				fmt.Println("failed to ask api key questions", err)
				return 1
			}

			if answers.PublicAPIKey != "" {
				cmd.flags.publicAPIKey = answers.PublicAPIKey
			}
			if answers.PrivateAPIKey != "" {
				cmd.flags.privateAPIKey = answers.PrivateAPIKey
			}
		}
	}
	return 0
}

func (cmd *createCommand) Help() string {
	return "user create help"
}

func (cmd *createCommand) Synopsis() string {
	return "user create synopsis"
}

type Flags struct {
	*flag.FlagSet

	username      string
	password      string
	publicAPIKey  string
	privateAPIKey string
}

func NewFlags(name string) Flags {
	var flags Flags

	flagSet := flag.NewFlagSet(name, flag.ContinueOnError)
	flagSet.StringVar(&flags.username, "username", "", "username")
	flagSet.StringVar(&flags.password, "password", "", "password")
	flagSet.StringVar(&flags.publicAPIKey, "publicAPIKey", "", "public api key")
	flagSet.StringVar(&flags.privateAPIKey, "privateAPIKey", "", "private api key")

	flags.FlagSet = flagSet
	return flags
}
