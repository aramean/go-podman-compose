package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

const message = "\nRun 'podman-compose COMMAND --help' for more information on a command"

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
	fmt.Println("  up          Create and start containers")
	fmt.Println("  down        Stop and remove containers, networks")
	fmt.Println("  exec        Execute a command in a running container")
	fmt.Println("  kill        Force stop service containers")
	fmt.Println("  logs        View output from containers")
	fmt.Println("  ps          List containers")
	fmt.Println("  pull        Pull service images")
	fmt.Println("  restart     Restart containers")
	fmt.Println("  start       Start services")
	fmt.Println("  stop        Stop services")
	fmt.Println("  version     Show Podman-Compose version information")
	return errors.New(message)
}

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

func runUsage() error {

	var arg0, arg1 = "", ""

	for i, v := range args {
		if i == 0 {
			arg0 = v
		}
		if i == 1 {
			arg1 = v
		}
	}

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

	return fmt.Errorf("%s\nunknown command: \"%s\"", message, arg0)
}
