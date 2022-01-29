package main

import (
	"flag"
	"strings"
)

type UpCommand struct {
	fs *flag.FlagSet

	name          string
	detach        bool
	removeOrphans bool
}

func NewUpCommand(args []string) *UpCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &UpCommand{
		fs: flag.NewFlagSet("up", exit),
	}

	gc.fs.BoolVar(&gc.detach, "d", false, "Detached mode: Run containers in the background")
	gc.fs.BoolVar(&gc.removeOrphans, "remove-orphans", false, "Remove containers for services not defined in the Compose file.")
	return gc
}

func (g *UpCommand) Name() string {
	return g.fs.Name()
}

func (g *UpCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *UpCommand) Run() error {
	//fmt.Println("Detach: ", g.detach, "!")
	return nil
}
