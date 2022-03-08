package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	messageError     = "\nRun 'podman-compose COMMAND --help' for more information on a command"
	messageUsage     = "Description:\n   %s\n\nUsage:\n  " + binaryName + " %s [options]\n\nOptions:\n"
	mainMessageUsage = "Usage:\n  " + binaryName + " [options] [command]\n\nAvailable Commands:\n"
)

type MainCommand struct {
	fs *flag.FlagSet

	detach        bool
	removeOrphans bool
}

func NewMainCommand() *MainCommand {

	gc := &MainCommand{
		fs: flag.NewFlagSet("help", flag.ExitOnError),
	}

	gc.fs.BoolVar(&gc.detach, "d", false, "Detached mode: Run containers in the background")
	gc.fs.BoolVar(&gc.removeOrphans, "remove-orphans", false, "Remove containers for services not defined in the Compose file.")
	return gc
}

func (g *MainCommand) Name() string {
	return g.fs.Name()
}

func (g *MainCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *MainCommand) Run() error {
	fmt.Print(mainMessageUsage)
	fmt.Fprintf(os.Stdout, "  build\t\t%s\n", DescriptionBuildCommand)
	fmt.Fprintf(os.Stdout, "  up\t\t%s\n", DescriptionUpCommand)
	fmt.Fprintf(os.Stdout, "  down\t\t%s\n", DescriptionDownCommand)
	fmt.Fprintf(os.Stdout, "  exec\t\t%s\n", DescriptionExecCommand)
	fmt.Fprintf(os.Stdout, "  kill\t\t%s\n", DescriptionKillCommand)
	fmt.Fprintf(os.Stdout, "  logs\t\t%s\n", DescriptionLogsCommand)
	fmt.Fprintf(os.Stdout, "  ps\t\t%s\n", DescriptionPsCommand)
	fmt.Fprintf(os.Stdout, "  pull\t\t%s\n", DescriptionPullCommand)
	fmt.Fprintf(os.Stdout, "  restart\t%s\n", DescriptionRestartCommand)
	fmt.Fprintf(os.Stdout, "  start\t\t%s\n", DescriptionStartCommand)
	fmt.Fprintf(os.Stdout, "  stop\t\t%s\n", DescriptionStopCommand)
	fmt.Fprintf(os.Stdout, "  version \t%s\n", DescriptionVersionCommand)
	return errors.New(messageError)
}

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func runUsage() error {
	if arg1 == "" {
		if arg0 == "" || strings.Contains(arg0, "-help") {
			args = append(args, "help")
			arg0 = "help"
		} else if arg0 == "-v" || arg0 == "--version" {
			arg0 = "version"
		}
	}

	if arg1 != "" {
		if strings.Contains(arg0, "-env-file") {
			arg0 = "version"
		}
	}

	cmds := []Runner{
		NewMainCommand(),
		NewBuildCommand(),
		NewDownCommand(),
		NewExecCommand(),
		NewKillCommand(),
		NewLogsCommand(),
		NewPsCommand(),
		NewPullCommand(),
		NewRestartCommand(),
		NewStartCommand(),
		NewStopCommand(),
		NewUpCommand(),
		NewVersionCommand(),
	}

	for _, cmd := range cmds {
		if cmd.Name() == arg0 {
			cmd.Init(args[1:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("%s\nunknown command: \"%s\"", messageError, arg0)
}

func displayUsage(fs *flag.FlagSet, out io.Writer, desc string, arg string) func() {
	fs.SetOutput(out)
	return func() {
		fmt.Fprintf(out, messageUsage, desc, arg)
		fs.PrintDefaults()
	}
}
