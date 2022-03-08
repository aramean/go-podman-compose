package main

import (
	"flag"
	"os"
)

const (
	DescriptionBuildCommand = "Build or rebuild services"
)

type BuildCommand struct {
	fs *flag.FlagSet
}

func NewBuildCommand() *BuildCommand {
	gc := &BuildCommand{
		fs: flag.NewFlagSet("build", flag.ExitOnError),
	}

	gc.fs.Usage = displayUsage(gc.fs, os.Stdout, DescriptionBuildCommand, "build")
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
