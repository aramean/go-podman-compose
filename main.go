package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	debug         = os.Getenv("DEBUG")
	binaryName    = "podman-compose"
	binaryVersion = "1.0.11"
	args          = os.Args[1:]
	detach        bool
	timeout       string
	removeOrphans bool
	signal        string
	tty           bool
	user          string
	quiet         bool
)

const (
	fileYML          = "docker-compose.yml"
	fileYAML         = "docker-compose.yaml"
	fileYMLOverride  = "docker-compose.override.yml"
	fileYAMLOverride = "docker-compose.override.yaml"
	fileEnv          = ".env"
)

const (
	colorBlue   = "\033[1;34m%s\033[0m"
	colorCyan   = "\033[1;36m%s\033[0m"
	colorGreen  = "\033[1;32m%s\033[0m"
	colorPurple = "\033[1;35m%s\033[0m"
	colorRed    = "\033[1;31m%s\033[0m"
	colorYellow = "\033[1;33m%s\033[0m"
	colorWhite  = "\033[1;37m%s\033[0m"
)

var arg0, arg1 string

type Command struct {
	OutputStatus   bool
	OutputNewlines bool
	Tasks          []CommandTask
}

func init() {
	if err := runUsage(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	for i, v := range args {
		if i == 0 {
			arg0 = v
		}
		if i == 1 {
			arg1 = v
		}
	}

	for i, f := range args {
		if f == "-d" || f == "--detach" {
			detach = true
		}
		if f == "-t" || f == "--timeout" {
			timeout = args[(i + 1)]
		}
		if f == "--remove-orphans" {
			removeOrphans = true
		}
		if f == "-q" || f == "--quiet" {
			quiet = true
		}
		if f == "-s" || f == "--signal" {
			signal = args[(i + 1)]
		}
		if f == "-u" || f == "--user" {
			user = args[(i + 1)]
		}
	}
}

func main() {
	l := loadEnvironmentVariables()
	e := parseYAML(l.Environment)
	g := buildCommand(e, l.Environment)

	for _, field := range g.Tasks {

		e := executeCommand(
			field.Command,
			field.OutputCustomMessage,
			field.OutputSingleline,
			field.OutputQuiet,
		)

		output := strings.TrimSpace(e.OutputCustomMessage)

		switch e.OutputStatusCode {
		case 0:
			if !field.OutputQuiet {
				if g.OutputStatus {
					fmt.Printf(colorGreen, "[OK] ")
				}
				if g.OutputNewlines {
					fmt.Println(output)
				} else {
					fmt.Print(output)
				}
			}
		case 1:

			m := strings.Split(output, "\n")

			for i, s := range m {
				if i == 0 {
					if s != "" {
						fmt.Printf(colorRed, "[**] ")
						fmt.Println(s)
					}
				} else {
					fmt.Printf(colorRed, "|**| "+s+"\n")
				}
			}
		}
	}
}

func buildCommand(e *Yaml, l []EnvironmentVariable) Command {

	g := Command{}

	switch arg0 {
	case "build":
		return newComposeBuildCommand(e, l)
	case "up", "start":
		return newComposeUpStartCommand(e, l)
	case "down":
		return newComposeDownCommand(e, l)
	case "exec":
		return newComposeExecCommand(e, l)
	case "kill":
		return newComposeKillCommand(e, l)
	case "restart":
		return newComposeRestartCommand(e, l)
	case "stop":
		return newComposeStopCommand(e, l)
	case "ps":
		return newComposePsCommand(e, l)
	case "pull":
		return newComposePullCommand(e, l)
	case "logs":
		return newComposeLogsCommand(e, l)
	}

	return g
}

func getContainerName(v Services, name string) string {
	if v.ContainerName != "" {
		name = v.ContainerName
	}
	return name
}
