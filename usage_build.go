package main

import (
	"flag"
)

type BuildCommand struct {
	fs *flag.FlagSet
}

func NewBuildCommand() *BuildCommand {
	gc := &BuildCommand{
		fs: flag.NewFlagSet("build", flag.ExitOnError),
	}

	return gc
}

func (g *BuildCommand) Name() string {
	return g.fs.Name()
}

func (g *BuildCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *BuildCommand) Run() error {
	return nil
}
