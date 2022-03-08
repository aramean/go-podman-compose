package main

import (
	"flag"
	"os"
)

const (
	DescriptionKillCommand = "Force stop service containers"
)

type KillCommand struct {
	fs *flag.FlagSet
}

func NewKillCommand() *KillCommand {
	gc := &KillCommand{
		fs: flag.NewFlagSet("kill", flag.ExitOnError),
	}

	gc.fs.Usage = displayUsage(gc.fs, os.Stdout, DescriptionKillCommand, "kill")
	return gc
}

func (g *KillCommand) Name() string {
	return g.fs.Name()
}

func (g *KillCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *KillCommand) Run() error {
	return nil
}
