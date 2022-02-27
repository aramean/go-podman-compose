package main

import (
	"flag"
)

type ExecCommand struct {
	fs *flag.FlagSet

	detach bool
	tty    bool
	user   string
}

func NewExecCommand() *ExecCommand {
	gc := &ExecCommand{
		fs: flag.NewFlagSet("exec", flag.ExitOnError),
	}

	gc.fs.BoolVar(&gc.detach, "d", false, "Detached mode: Run command in the background.")
	gc.fs.BoolVar(&gc.tty, "T", false, "Disable pseudo-TTY allocation. By default "+binaryName+" exec allocates a TTY.")
	gc.fs.StringVar(&gc.user, "u", "", "Run the command as this user.")

	return gc
}

func (g *ExecCommand) Name() string {
	return g.fs.Name()
}

func (g *ExecCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *ExecCommand) Run() error {
	return nil
}
