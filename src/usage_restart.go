package main

import (
	"flag"
)

type RestartCommand struct {
	fs *flag.FlagSet

	timeout int
}

func NewRestartCommand() *RestartCommand {
	gc := &RestartCommand{
		fs: flag.NewFlagSet("restart", flag.ExitOnError),
	}

	gc.fs.IntVar(&gc.timeout, "t", 10, "Specify a shutdown timeout in seconds (default 10)")
	return gc
}

func (g *RestartCommand) Name() string {
	return g.fs.Name()
}

func (g *RestartCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *RestartCommand) Run() error {
	return nil
}
