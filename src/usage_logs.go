package main

import (
	"flag"
	"strings"
)

type LogsCommand struct {
	fs *flag.FlagSet
}

func NewLogsCommand(args []string) *LogsCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &LogsCommand{
		fs: flag.NewFlagSet("logs", exit),
	}

	return gc
}

func (g *LogsCommand) Name() string {
	return g.fs.Name()
}

func (g *LogsCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *LogsCommand) Run() error {
	return nil
}
