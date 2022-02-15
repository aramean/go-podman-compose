package main

import (
	"fmt"

	"github.com/go-cmd/cmd"
)

const binary = "podman"

type CommandTask struct {
	Command             []string
	OutputCustomMessage string
	OutputStatusCode    int
	OutputSingleline    bool
	OutputQuiet         bool
}

func executeCommand(f []string, m string, s bool, q bool) *CommandTask {

	if debug == "TRUE" {
		debugMessage := fmt.Sprintln(f)
		fmt.Printf(colorYellow, debugMessage)
	}

	var customMessage = len(m)

	t := CommandTask{
		Command:             []string{},
		OutputCustomMessage: m,
		OutputSingleline:    s,
		OutputQuiet:         q,
	}

	// Disable output buffering, enable streaming
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	// Create Cmd with options
	envCmd := cmd.NewCmdOptions(cmdOptions, binary, f...)

	// Print STDOUT and STDERR lines streaming from Cmd
	doneChan := make(chan struct{})
	go func() {
		defer close(doneChan)
		// Done when both channels have been closed
		// https://dave.cheney.net/2013/04/30/curious-channels
		for envCmd.Stdout != nil || envCmd.Stderr != nil {
			select {
			case line, open := <-envCmd.Stdout:
				if !open {
					envCmd.Stdout = nil
					continue
				}

				if customMessage == 0 && !t.OutputQuiet {
					t.OutputCustomMessage = t.OutputCustomMessage + "\n" + line
				}
				t.OutputStatusCode = 0

			case line, open := <-envCmd.Stderr:
				if !open {
					envCmd.Stderr = nil
					continue
				}

				if !t.OutputQuiet {
					t.OutputCustomMessage = t.OutputCustomMessage + "\n" + line
				}
				t.OutputStatusCode = 1
			}
		}
	}()

	// Run and wait for Cmd to return, discard Status
	<-envCmd.Start()

	// Wait for goroutine to print everything
	<-doneChan

	return &t
}
