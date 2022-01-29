package main

import (
	"flag"
	"strings"
)

type DownCommand struct {
	fs *flag.FlagSet

	removeOrphans bool
}

func NewDownCommand(args []string) *DownCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &DownCommand{
		fs: flag.NewFlagSet("down", exit),
	}

	gc.fs.BoolVar(&gc.removeOrphans, "remove-orphans", false, "Remove containers for services not defined in the Compose file.")
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
