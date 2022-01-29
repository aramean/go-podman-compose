package main

import (
	"flag"
	"strings"
)

type StopCommand struct {
	fs *flag.FlagSet

	timeout int
}

func NewStopCommand(args []string) *StopCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &StopCommand{
		fs: flag.NewFlagSet("stop", exit),
	}

	gc.fs.IntVar(&gc.timeout, "t", 10, "Specify a shutdown timeout in seconds (default 10)")
	return gc
}

func (g *StopCommand) Name() string {
	return g.fs.Name()
}

func (g *StopCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *StopCommand) Run() error {
	return nil
}
