package common

import (
	"flag"
	"fmt"
)

type Command struct {
	Meta     CommandMeta
	Preparer CommandPreparer
	Runner   CommandRunner
	Reporter CommandReporter
}

type CommandMeta struct {
	Name        string
	Description string
}

func (cmd *Command) Synopsis() string {
	return cmd.Meta.Description
}

func (cmd *Command) Help() string {
	return fmt.Sprintf("The %s command allows you to %s", cmd.Meta.Name, cmd.Meta.Description)
}

func (cmd *Command) Run(args []string) int {
	return 0
}

type CommandFlagger interface {
	NewFlagSet(name string) *flag.FlagSet
}

type CommandFlags struct {
	ConfigPath    string
	RealmBaseURL  string
	CloudBaseURL  string
	DisableColors bool
	OutputFormat  CommandOutputFormat
	OutputTarget  string
	AutoConfirm   bool
}

func (flags CommandFlags) NewFlagSet(name string) *flag.FlagSet {
	set := flag.NewFlagSet(name, flag.ContinueOnError)
	set.Usage = func() {}

	set.StringVar(&flags.ConfigPath, "config-path", "", "")
	set.StringVar(&flags.RealmBaseURL, "base-url", "https://realm.mongodb.org", "")
	set.StringVar(&flags.CloudBaseURL, "atlas-base-url", "https://cloud.mongodb.org", "")
	set.BoolVar(&flags.DisableColors, "disable-colors", false, "")
	set.StringVar(&flags.OutputTarget, "output-format", string(CommandOutputFormatText), "")
	set.StringVar(&flags.OutputTarget, "output-target", "", "")
	set.BoolVar(&flags.AutoConfirm, "yes", false, "")
	set.BoolVar(&flags.AutoConfirm, "y", false, "")

	return set
}

type CommandOutputFormat string

const (
	CommandOutputFormatText CommandOutputFormat = "text"
	CommandOutputFormatJSON CommandOutputFormat = "json"
)

type CommandPreparer interface {
	SetFlags() CommandFlagger
	CollectInputs() error
	ValidateInputs() error
	Setup() (*Context, UI, error)
}

type CommandRunner interface {
	Run(ctx *Context, ui UI) (int, CommandFeedback, error)
}

type CommandReporter interface {
	PrintCompletionMessage(exitCode int)
	PrintFeedback(CommandFeedback)
	PrintFollowupItems()
}

type CommandFeedback interface{}
