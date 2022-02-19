package main

import (
	"flag"
	"fmt"
)

type VersionCommand struct {
	fs *flag.FlagSet

	format string
}

func NewVersionCommand() *VersionCommand {
	gc := &VersionCommand{
		fs: flag.NewFlagSet("version", flag.ExitOnError),
	}

	gc.fs.StringVar(&gc.format, "f", "", "")
	return gc
}

func (g *VersionCommand) Name() string {
	return g.fs.Name()
}

func (g *VersionCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *VersionCommand) Run() error {
	fmt.Print(binaryName + " version " + binaryVersion)
	return fmt.Errorf("")
}
