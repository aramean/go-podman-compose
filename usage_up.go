package main

import (
	"flag"
	"os"
)

const (
	DescriptionUpCommand = "Create and start containers"
)

type UpCommand struct {
	fs *flag.FlagSet

	detach        bool
	removeOrphans bool
}

func NewUpCommand() *UpCommand {
	gc := &UpCommand{
		fs: flag.NewFlagSet("up", flag.ExitOnError),
	}

	gc.fs.BoolVar(&gc.detach, "d", false, "Detached mode: Run containers in the background")
	gc.fs.BoolVar(&gc.removeOrphans, "remove-orphans", false, "Remove containers for services not defined in the Compose file.")
	gc.fs.Usage = displayUsage(gc.fs, os.Stdout, DescriptionUpCommand, "up")
	return gc
}

func (g *UpCommand) Name() string {
	return g.fs.Name()
}

func (g *UpCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *UpCommand) Run() error {
	return nil
}
