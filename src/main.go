package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-cmd/cmd"
	"gopkg.in/yaml.v3"
)

const (
	ColorInfo    = "\033[1;34m%s\033[0m"
	ColorNotice  = "\033[1;36m%s\033[0m"
	ColorWarning = "\033[1;33m%s\033[0m"
	ColorError   = "\033[1;31m%s\033[0m"
	ColorDebug   = "\033[0;36m%s\033[0m"
)

var (
	version *string
	debug   *bool
)

type Services struct {
	Image   string   `yaml:"image"`
	Volumes []string `yaml:"volumes"`
	Ports   []string `yaml:"ports"`
}

type Volumes struct {
	Mounts []string `yaml:"volumes"`
}

type Networks struct {
	Networks []string `yaml:"networks"`
}

type Config struct {
	SkipHeaderValidation bool
	Services             map[string]Services
	Volumes              map[string]Volumes
}

func (e *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	/*var params struct {
		SkipHeaderValidation bool `yaml:"skip-header-validation"`
	}
	if err := unmarshal(&params); err != nil {
		return err
	}*/
	var services map[string]Services
	if err := unmarshal(&services); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}
	e.Services = services

	var volumes map[string]Volumes
	//e.SkipHeaderValidation = params.SkipHeaderValidation
	e.Volumes = volumes
	return nil
}

func main() {
	var arg1 = os.Args[1:]

	if err := root(arg1); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	e := parseYML()
	f := buildCommand(e, arg1[0])

	for i := range f {
		//fmt.Printf("Row: %v\n", i)
		//fmt.Println(f[i])
		executeCommand(f[i])
	}
}

func parseYML() map[string]Config {
	yfile, err2 := ioutil.ReadFile("docker-compose.yml")

	if err2 != nil {
		log.Fatal(err2)
	}

	var e map[string]Config
	err3 := yaml.Unmarshal(yfile, &e)

	if err3 != nil {
		log.Fatal(err3)
	}

	return e
}

func buildCommand(e map[string]Config, arg1 string) [][]string {

	args := [][]string{}

	switch arg1 {
	case "ps":
		arr := []string{"ps"}
		args = append(args, arr)
		break
	case "up":
		for _, v := range e {
			for k, v := range v.Services {
				arr := []string{"run", "--replace", "-d", "--name", k, "-v", v.Volumes[0], "-p", v.Ports[0], v.Image}
				args = append(args, arr)
			}
		}
		break
	case "down":
		for _, v := range e {
			for k, _ := range v.Services {
				arr := []string{"stop", "--time", "2", k}
				args = append(args, arr)
			}
		}
		break
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

	if message != "" {
		//fmt.Printf("%s %s ... ", message, name)
		fmt.Printf(ColorInfo, " ... OK\n")
		//fmt.Printf("\n\n%#v", args)
	}
}

func runCommand(arg, name string, image string, volumes []string, ports []string) {

	args := []string{}
	message := ""

	switch arg {
	case "ps":
		args = []string{"ps"}
		break
	case "stop":
		message = "Stopping"
		args = []string{"stop", name}
		break
	case "start":
		message = "Starting"
		args = []string{"start", name}
		break
	case "up":
		message = "Starting"
		args = []string{"run", "--replace", "-d", "--name", name, "-v", volumes[0], "-p", ports[0], image}
		break
	case "down":
		message = "Stopping"
		args = []string{"stop", "--time", "2", name}
		break
	}

	// Disable output buffering, enable streaming
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}

	// Create Cmd with options
	envCmd := cmd.NewCmdOptions(cmdOptions, "podman", args...)

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

	if message != "" {
		//fmt.Printf("%s %s ... ", message, name)
		fmt.Printf(ColorInfo, " ... OK\n")
		//fmt.Printf("\n\n%#v", args)
	}
}
