package main

import (
	"flag"
	"os"
)

const (
	DescriptionDownCommand = "Stop and remove containers, networks"
)

type DownCommand struct {
	fs *flag.FlagSet

	removeOrphans bool
}

func NewDownCommand() *DownCommand {
	gc := &DownCommand{
		fs: flag.NewFlagSet("down", flag.ExitOnError),
	}

	gc.fs.BoolVar(&gc.removeOrphans, "remove-orphans", false, "Remove containers for services not defined in the Compose file.")
	gc.fs.Usage = displayUsage(gc.fs, os.Stdout, DescriptionDownCommand, "down")
	return gc
}

func (g *DownCommand) Name() string {
	return g.fs.Name()
}

func (g *DownCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *DownCommand) Run() error {
	return nil
}
