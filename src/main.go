package main

import (
	"fmt"
	"os"

	"github.com/go-cmd/cmd"
)

const (
	ColorInfo    = "\033[1;34m%s\033[0m"
	ColorNotice  = "\033[1;36m%s\033[0m"
	ColorWarning = "\033[1;33m%s\033[0m"
	ColorError   = "\033[1;31m%s\033[0m"
	ColorDebug   = "\033[0;36m%s\033[0m"
)

func main() {
	var arg = os.Args[1:]

	if err := root(arg); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	e := parseYML()
	f := buildCommand(e, arg)

	for i := range f {
		/*fmt.Printf("Row: %v\n", i)
		fmt.Println(f[i])
		*/
		executeCommand(f[i])

		if message != "" {
			//fmt.Printf("%s %s ... ", message, name)
			fmt.Printf(ColorInfo, " ... \n")
			//fmt.Printf("\n\n%#v", args)
		}
	}
}

func buildCommand(e map[string]Config, arg []string) [][]string {

	var arg0, arg1 = "", ""

	for i := range arg {
		if i == 0 {
			arg0 = arg[0]
		}
		if i == 1 {
			arg1 = arg[1]
		}
	}

	args := [][]string{}

	switch arg0 {
	case "down":
		for _, v := range e {
			for k, _ := range v.Services {
				arr1 := []string{"stop", "--time", "2", k}
				arr2 := []string{"rm", k}
				args = append(args, arr1, arr2)
			}
		}
		return args
	case "ps":
		arr := []string{"ps"}
		args = append(args, arr)
		return args
	case "start":
		for _, v := range e {
			for k, v := range v.Services {
				if (len(arg1) > 0 && k == arg1) || len(arg1) == 0 {
					arr := []string{
						"run",
						"--replace",
						"-d",
						"--name", k,
					}

					for i := range v.Ports {
						arr = append(arr, "-p", v.Ports[i])
					}

					for i := range v.Volumes {
						arr = append(arr, "-v", v.Volumes[i])
					}

					for i := range v.Environment {
						arr = append(arr, "-e", v.Environment[i])
					}

					arr = append(arr, v.Image)
					args = append(args, arr)
				}
			}
		}
		return args
	case "stop":
		for _, v := range e {
			for k, _ := range v.Services {
				if (len(arg1) > 0 && k == arg1) || len(arg1) == 0 {
					arr := []string{
						"stop",
						"--time", "2",
						k}
					args = append(args, arr)
				}
			}
		}
		return args
	case "up":
		for _, v := range e {
			for k, v := range v.Services {

				arr := []string{
					"run",
					"--replace",
					"-d",
					"--name", k,
				}

				for i := range v.Ports {
					arr = append(arr, "-p", v.Ports[i])
				}

				for i := range v.Volumes {
					arr = append(arr, "-v", v.Volumes[i])
				}

				for i := range v.Environment {
					arr = append(arr, "-e", v.Environment[i])
				}

				arr = append(arr, v.Image)
				args = append(args, arr)
			}
		}
		return args
	}

	return args
}

func executeCommand(f []string) {
	// Disable output buffering, enable streaming
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	// Create Cmd with options
	envCmd := cmd.NewCmdOptions(cmdOptions, "podman", f...)

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
				fmt.Print(line)
			case line, open := <-envCmd.Stderr:
				if !open {
					envCmd.Stderr = nil
					continue
				}
				fmt.Fprintln(os.Stderr, line)
			}
		}
	}()

	// Run and wait for Cmd to return, discard Status
	<-envCmd.Start()

	// Wait for goroutine to print everything
	<-doneChan
}
