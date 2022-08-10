package common

import (
	"os"

	"github.com/mitchellh/cli"
)

func NewUI() cli.Ui {
	return &cli.ConcurrentUi{Ui: NewBasicUI()}
}

func NewBasicUI() cli.Ui {
	return &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}
}
