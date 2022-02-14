package main

import (
	"flag"
)

type PsCommand struct {
	fs *flag.FlagSet

	all      string
	external string
}

func NewPsCommand() *PsCommand {
	gc := &PsCommand{
		fs: flag.NewFlagSet("ps", flag.ExitOnError),
	}

	gc.fs.StringVar(&gc.all, "a", "all", "Show all the containers, default is only running containers")
	gc.fs.StringVar(&gc.external, "external", "", "Show containers in storage not controlled by Podman")
	return gc
}

func (g *PsCommand) Name() string {
	return g.fs.Name()
}

func (g *PsCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *PsCommand) Run() error {
	return nil
}
