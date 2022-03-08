package main

import (
	"flag"
	"os"
)

const (
	DescriptionPullCommand = "Pull service images"
)

type PullCommand struct {
	fs *flag.FlagSet
}

func NewPullCommand() *PullCommand {
	gc := &PullCommand{
		fs: flag.NewFlagSet("pull", flag.ExitOnError),
	}

	gc.fs.Usage = displayUsage(gc.fs, os.Stdout, DescriptionPullCommand, "pull")
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
