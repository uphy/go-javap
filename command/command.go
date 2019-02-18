package command

import (
	"github.com/urfave/cli"
)

type CLI struct {
	app *cli.App
}

func New() *CLI {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		listCommand(),
	}
	return &CLI{app}
}

func (c *CLI) Execute(args []string) error {
	return c.app.Run(args)
}
