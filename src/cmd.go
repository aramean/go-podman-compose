package main

import (
	"github.com/go-cmd/cmd"
)

const binary = "podman"

type cmdMessage struct {
	statusCode int
	message    string
	newlines   bool
}

func executeCommand(f []string) *cmdMessage {

	m := cmdMessage{
		statusCode: 0,
		message:    "",
		newlines:   false,
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

				m.statusCode = 0
				m.message = m.message + line + " "

			case line, open := <-envCmd.Stderr:
				if !open {
					envCmd.Stderr = nil
					continue
				}

				m.statusCode = 1
				m.message = m.message + line + " "
			}

			/*debugMessage := fmt.Sprintln(envCmd.Args)
			fmt.Printf(colorYellow, debugMessage)*/
		}
	}()

	// Run and wait for Cmd to return, discard Status
	<-envCmd.Start()

	// Wait for goroutine to print everything
	<-doneChan

	return &m
}
