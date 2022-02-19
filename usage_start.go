package main

import (
	"flag"
)

type StartCommand struct {
	fs *flag.FlagSet
}

func NewStartCommand() *StartCommand {
	gc := &StartCommand{
		fs: flag.NewFlagSet("start", flag.ExitOnError),
	}

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
