package main

import (
	"flag"
)

type StopCommand struct {
	fs *flag.FlagSet

	timeout int
}

func NewStopCommand() *StopCommand {
	gc := &StopCommand{
		fs: flag.NewFlagSet("stop", flag.ExitOnError),
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
