package main

import (
	"flag"
	"os"
)

const (
	DescriptionLogsCommand = "View output from containers"
)

type LogsCommand struct {
	fs *flag.FlagSet
}

func NewLogsCommand() *LogsCommand {
	gc := &LogsCommand{
		fs: flag.NewFlagSet("logs", flag.ExitOnError),
	}

	gc.fs.Usage = displayUsage(gc.fs, os.Stdout, DescriptionLogsCommand, "logs")
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
