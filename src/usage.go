package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

const message = "\nRun 'podman-compose COMMAND --help' for more information on a command."

type MainCommand struct {
	fs *flag.FlagSet

	name          string
	detach        bool
	removeOrphans bool
}

type DownCommand struct {
	fs *flag.FlagSet

	name          string
	detach        bool
	removeOrphans bool
}

type PsCommand struct {
	fs *flag.FlagSet

	all      string
	external string
}

type StopCommand struct {
	fs *flag.FlagSet

	timeout int
}

type UpCommand struct {
	fs *flag.FlagSet

	name          string
	detach        bool
	removeOrphans bool
}

type VersionCommand struct {
	fs *flag.FlagSet

	format string
}

func NewMainCommand(args []string) *MainCommand {

	gc := &MainCommand{
		fs: flag.NewFlagSet("help", flag.ExitOnError),
	}

	gc.fs.BoolVar(&gc.detach, "d", false, "Detached mode: Run containers in the background")
	gc.fs.BoolVar(&gc.removeOrphans, "remove-orphans", false, "Remove containers for services not defined in the Compose file.")
	return gc
}

func NewDownCommand(args []string) *UpCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &UpCommand{
		fs: flag.NewFlagSet("down", exit),
	}

	gc.fs.BoolVar(&gc.detach, "d", false, "Detached mode: Run containers in the background")
	gc.fs.BoolVar(&gc.removeOrphans, "remove-orphans", false, "Remove containers for services not defined in the Compose file.")
	return gc
}

func NewPsCommand(args []string) *PsCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &PsCommand{
		fs: flag.NewFlagSet("ps", exit),
	}

	gc.fs.StringVar(&gc.all, "a", "all", "Show all the containers, default is only running containers")
	gc.fs.StringVar(&gc.external, "external", "", "Show containers in storage not controlled by Podman")
	return gc
}

func NewStopCommand(args []string) *StopCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &StopCommand{
		fs: flag.NewFlagSet("stop", exit),
	}

	gc.fs.IntVar(&gc.timeout, "t", 10, "Specify a shutdown timeout in seconds (default 10)")
	return gc
}

func NewUpCommand(args []string) *UpCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &UpCommand{
		fs: flag.NewFlagSet("up", exit),
	}

	gc.fs.BoolVar(&gc.detach, "d", false, "Detached mode: Run containers in the background")
	gc.fs.BoolVar(&gc.removeOrphans, "remove-orphans", false, "Remove containers for services not defined in the Compose file.")
	return gc
}

func NewVersionCommand(args []string) *VersionCommand {

	var exit = flag.ContinueOnError
	if len(args) > 1 && strings.Contains(args[1], "-help") {
		exit = flag.ExitOnError
	}

	gc := &VersionCommand{
		fs: flag.NewFlagSet("version", exit),
	}

	gc.fs.StringVar(&gc.format, "f", "", "")
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
	fmt.Println("  version     Show the Podman-Compose version information")
	return errors.New(message)
}

func (g *DownCommand) Name() string {
	return g.fs.Name()
}

func (g *DownCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *DownCommand) Run() error {
	fmt.Println("Detach: ", g.detach, "!")
	return nil
}

func (g *PsCommand) Name() string {
	return g.fs.Name()
}

func (g *PsCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *PsCommand) Run() error {
	return nil
}

func (g *StopCommand) Name() string {
	return g.fs.Name()
}

func (g *StopCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *StopCommand) Run() error {
	return nil
}

func (g *UpCommand) Name() string {
	return g.fs.Name()
}

func (g *UpCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *UpCommand) Run() error {
	//fmt.Println("Detach: ", g.detach, "!")
	return nil
}

func (g *VersionCommand) Name() string {
	return g.fs.Name()
}

func (g *VersionCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *VersionCommand) Run() error {
	fmt.Print("Podman-Compose version v1.0.0")
	return fmt.Errorf("")
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
