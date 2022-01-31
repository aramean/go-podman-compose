package main

import (
	"flag"
	"strings"
)

type RestartCommand struct {
	fs *flag.FlagSet
}

func NewRestartCommand(args []string) *RestartCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &RestartCommand{
		fs: flag.NewFlagSet("restart", exit),
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
