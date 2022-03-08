package main

import (
	"flag"
	"os"
)

const (
	DescriptionPsCommand = "List containers"
)

type PsCommand struct {
	fs *flag.FlagSet

	all      string
	external string
	quiet    bool
}

func NewPsCommand() *PsCommand {
	gc := &PsCommand{
		fs: flag.NewFlagSet("ps", flag.ExitOnError),
	}

	gc.fs.StringVar(&gc.all, "a", "all", "Show all the containers, default is only running containers")
	gc.fs.StringVar(&gc.external, "external", "", "Show containers in storage not controlled by Podman")
	gc.fs.BoolVar(&gc.quiet, "q", false, "Print the numeric IDs of the containers only")
	gc.fs.Usage = displayUsage(gc.fs, os.Stdout, DescriptionPsCommand, "ps")
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
