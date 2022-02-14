package main

import (
	"flag"
	"strings"
)

type KillCommand struct {
	fs *flag.FlagSet
}

func NewKillCommand(args []string) *KillCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &KillCommand{
		fs: flag.NewFlagSet("kill", exit),
	}

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
