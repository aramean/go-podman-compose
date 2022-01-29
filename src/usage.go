package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
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

func root(args []string) error {
	if len(args) < 1 {
		args[0] = "help"
		fmt.Println(args)
		os.Exit(0)
	}

	cmds := []Runner{
		NewMainCommand(args),
		NewDownCommand(args),
		NewPsCommand(args),
		NewStopCommand(args),
		NewUpCommand(args),
		NewVersionCommand(args),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("%s\nunknown command: \"%s\"", message, subcommand)
}
