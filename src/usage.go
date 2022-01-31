package main

import (
	"errors"
	"flag"
	"fmt"
)

const message = "\nRun 'podman-compose COMMAND --help' for more information on a command."

type MainCommand struct {
	fs *flag.FlagSet

	detach        bool
	removeOrphans bool
}

func NewMainCommand(args []string) *MainCommand {

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
	fmt.Println("  ps          List containers")
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

func runUsage(args []string) error {

	if len(args) < 1 {
		args = append(args, "help")
	}

	cmds := []Runner{
		NewMainCommand(args),
		NewDownCommand(args),
		NewPsCommand(args),
		NewRestartCommand(args),
		NewStartCommand(args),
		NewStopCommand(args),
		NewUpCommand(args),
		NewVersionCommand(args),
	}

	subcommand := args[0]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(args[1:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("%s\nunknown command: \"%s\"", message, subcommand)
}
