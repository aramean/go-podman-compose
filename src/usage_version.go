package main

import (
	"flag"
	"fmt"
	"strings"
)

type VersionCommand struct {
	fs *flag.FlagSet

	format string
}

func NewVersionCommand(args []string) *VersionCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &VersionCommand{
		fs: flag.NewFlagSet("version", exit),
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
	fmt.Print("Podman-Compose version v1.0.0")
	return fmt.Errorf("")
}
