package main

import (
	"flag"
	"os"
)

const (
	DescriptionRestartCommand = "Restart containers"
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
	gc.fs.Usage = displayUsage(gc.fs, os.Stdout, DescriptionRestartCommand, "restart")
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
