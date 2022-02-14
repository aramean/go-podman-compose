package main

import (
	"fmt"
	"os"
	"strings"
)

var debug = os.Getenv("DEBUG")

const (
	fileYAML         = "docker-compose.yml"
	fileYAMLOverride = "docker-compose.override.yml"
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

type Command struct {
	OutputStatus   bool
	OutputNewlines bool
	Tasks          []CommandTask
}

func main() {
	var arg = os.Args[1:]

	if err := runUsage(arg); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	l := loadEnvironmentVariables()
	e := parseYAML(l.Environment)

	g := buildCommand(e, l.Environment, arg)

	for _, field := range g.Tasks {

		e := executeCommand(
			field.Command,
			field.OutputMessage,
			field.OutputSingleline,
			field.OutputQuiet,
		)

		switch e.OutputStatusCode {
		case 0:
			if !field.OutputQuiet {
				if g.OutputStatus {
					fmt.Printf(colorGreen, "[OK] ")
				}
				if g.OutputNewlines {
					fmt.Println(e.OutputMessage)
				} else {
					fmt.Print(e.OutputMessage)
				}
			}
		case 1:

			m := strings.Split(e.OutputMessage, "\n")

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

func buildCommand(e *Config, l []EnvironmentVariable, arg []string) Command {

	var arg0, arg1 = "", ""

	for i, v := range arg {
		if i == 0 {
			arg0 = v
		}
		if i == 1 {
			arg1 = v
		}
	}

	g := Command{}

	switch arg0 {
	case "up":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k, v := range e.Services {

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

			for _, r := range convertEnvironmentVariable(v.Environment) {
				r := transformEnvironmentVariable(r, l)
				arr = append(arr, "-e", r.Name+"="+r.Value)
			}

			if v.Restart != "" {
				arr = append(arr, "--restart", v.Restart)
			}

			arr = append(arr, v.Image)

			g.Tasks = append(
				g.Tasks,
				CommandTask{arr, "Starting container " + k + " ...", 0, false, false},
			)

		}

	case "down":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k, _ := range e.Services {
			g.Tasks = append(
				g.Tasks,
				CommandTask{[]string{"stop", "--time", "2", k}, "Stopping container " + k + " ...", 0, true, false},
				CommandTask{[]string{"rm", k}, "", 0, true, true},
			)
		}
	case "kill":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k := range e.Services {

			arr := []string{"kill"}

			if len(arg1) > 0 {
				arr = append(arr, arg1)
			} else {
				arr = append(arr, "--all")
			}

			g.Tasks = append(
				g.Tasks,
				CommandTask{arr, "Killing container " + k + " ...", 0, true, false},
			)

		}
	case "restart":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k := range e.Services {
			if (len(arg1) > 0 && k == arg1) || len(arg1) == 0 {
				g.Tasks = append(
					g.Tasks,
					CommandTask{[]string{"restart", k}, "Restarting container " + k + " ...", 0, true, false},
				)
			}
		}
	case "stop":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k := range e.Services {
			if (len(arg1) > 0 && k == arg1) || len(arg1) == 0 {
				g.Tasks = append(
					g.Tasks,
					CommandTask{[]string{"stop", "--time", "2", k}, "Stopping container " + k + " ...", 0, true, false},
				)
			}
		}
	case "start":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k, v := range e.Services {

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

				for _, r := range convertEnvironmentVariable(v.Environment) {
					r := transformEnvironmentVariable(r, l)
					arr = append(arr, "-e", r.Name+"="+r.Value)
				}

				if v.Restart != "" {
					arr = append(arr, "--restart", v.Restart)
				}

				arr = append(arr, v.Image)

				g.Tasks = append(
					g.Tasks,
					CommandTask{arr, "Starting container " + k + " ...", 0, true, false},
				)
			}

		}
	case "ps":
		g = Command{
			OutputStatus: false,
			Tasks: []CommandTask{
				{
					[]string{"ps"}, "", 0, true, false,
				},
			},
		}
	case "logs":
		g = Command{
			OutputStatus:   false,
			OutputNewlines: true,
		}

		for k := range e.Services {
			if (len(arg1) > 0 && k == arg1) || len(arg1) == 0 {
				g.Tasks = append(
					g.Tasks,
					CommandTask{[]string{"logs", k}, "", 0, true, false},
				)
			}
		}
	}

	return g
}
