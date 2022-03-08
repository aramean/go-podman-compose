package main

import (
	"flag"
	"os"
)

const (
	DescriptionStartCommand = "Start services"
)

type StartCommand struct {
	fs *flag.FlagSet
}

func NewStartCommand() *StartCommand {
	gc := &StartCommand{
		fs: flag.NewFlagSet("start", flag.ExitOnError),
	}

	gc.fs.Usage = displayUsage(gc.fs, os.Stdout, DescriptionStartCommand, "start")
	return gc
}

func (g *StartCommand) Name() string {
	return g.fs.Name()
}

func (g *StartCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *StartCommand) Run() error {
	return nil
}
