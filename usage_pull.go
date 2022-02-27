package main

import (
	"flag"
)

type PullCommand struct {
	fs *flag.FlagSet
}

func NewPullCommand() *PullCommand {
	gc := &PullCommand{
		fs: flag.NewFlagSet("pull", flag.ExitOnError),
	}

	return gc
}

func (g *PullCommand) Name() string {
	return g.fs.Name()
}

func (g *PullCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *PullCommand) Run() error {
	return nil
}
