package main

import (
	"flag"
	"fmt"
	"strings"
)

type DownCommand struct {
	fs *flag.FlagSet

	name          string
	detach        bool
	removeOrphans bool
}

func NewDownCommand(args []string) *UpCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &UpCommand{
		fs: flag.NewFlagSet("down", exit),
	}

	gc.fs.BoolVar(&gc.detach, "d", false, "Detached mode: Run containers in the background")
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
	fmt.Println("Detach: ", g.detach, "!")
	return nil
}
