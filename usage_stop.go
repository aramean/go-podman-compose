package main

import (
	"flag"
	"os"
)

const (
	DescriptionStopCommand = "Stop services"
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
	gc.fs.Usage = displayUsage(gc.fs, os.Stdout, DescriptionStopCommand, "stop")
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
