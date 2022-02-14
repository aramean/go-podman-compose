package main

import (
	"flag"
)

type RestartCommand struct {
	fs *flag.FlagSet
}

func NewRestartCommand() *RestartCommand {
	gc := &RestartCommand{
		fs: flag.NewFlagSet("restart", flag.ExitOnError),
	}

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
