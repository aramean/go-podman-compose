package main

import (
	"flag"
)

type LogsCommand struct {
	fs *flag.FlagSet
}

func NewLogsCommand() *LogsCommand {
	gc := &LogsCommand{
		fs: flag.NewFlagSet("logs", flag.ExitOnError),
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
