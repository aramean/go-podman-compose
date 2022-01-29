package main

import (
	"flag"
	"strings"
)

type StartCommand struct {
	fs *flag.FlagSet
}

func NewStartCommand(args []string) *StartCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &StartCommand{
		fs: flag.NewFlagSet("start", exit),
	}

	return gc
}

func (g *StartCommand) Name() string {
	return g.fs.Name()
}

func (g *StartCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *StartCommand) Run() error {
	return nil
}
