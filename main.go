package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	debug         = os.Getenv("DEBUG")
	binaryName    = "podman-compose"
	binaryVersion = "1.0.3"
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

func init() {
	if err := runUsage(); err != nil {
		fmt.Println(err)
		os.Exit(0)
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

func buildCommand(e *Config, l []EnvironmentVariable) Command {

	var arg0, arg1 = "", ""

	for i, v := range args {
		if i == 0 {
			arg0 = v
		}
		if i == 1 {
			arg1 = v
		}
	}

	g := Command{}

	switch arg0 {
	case "up", "start":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k, v := range e.Services {
			if (len(arg1) > 0 && k == arg1) || len(arg1) == 0 {
				arr := []string{
					"run",
					"--replace",
					"--name", k,
					"-d",
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

				if v.CpuShares != "" {
					arr = append(arr, "--cpu-shares", v.CpuShares)
				}

				if v.Pid != "" {
					arr = append(arr, "--pid", v.Pid)
				}

				if v.Platform != "" {
					arr = append(arr, "--platform", v.Platform)
				}

				if v.Restart != "" {
					arr = append(arr, "--restart", v.Restart)
				}

				if v.EnvFile != nil {
					for _, r := range convertToArray(v.EnvFile) {
						arr = append(arr, "--env-file", r)
					}
				}

				if v.Expose != nil {
					for _, r := range convertToArray(v.Expose) {
						arr = append(arr, "--expose", r)
					}
				}

				if v.Dns != nil {
					for _, r := range convertToArray(v.Dns) {
						arr = append(arr, "--dns", r)
					}
				}

				if v.Init != "" {
					if _, err := strconv.ParseBool(v.Init); err == nil {
						arr = append(arr, "--init")
					} else {
						arr = append(arr, "--init-path", v.Init)
					}
				}

				arr = append(arr, v.Image)

				g.Tasks = append(
					g.Tasks,
					CommandTask{arr, "Starting container " + k + " ...", 0, false, false},
				)
			}
		}

		/*if !detach {

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

		}*/

		for k := range e.Volumes {
			g.Tasks = append(
				g.Tasks,
				CommandTask{[]string{"volume", "create", k}, "", 0, true, true},
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
	case "exec":
		g = Command{
			OutputStatus:   false,
			OutputNewlines: true,
		}

		var exec = ""
		var service = ""

		for i, f := range args {
			for k := range e.Services {
				if f == k {
					service = args[i]
					exec = args[(i + 1)]

					fmt.Print("hej")
				}
			}
		}

		arr := []string{
			"exec",
			"-it",
		}

		/*if detach {
			arr = append(arr, "--detach")
		}

		if tty {
			arr = append(arr, "--tty")
		}

		if len(user) > 0 {
			arr = append(arr, "--user", "root")
		}
		*/
		arr = append(arr, service, exec)

		g.Tasks = append(
			g.Tasks,
			CommandTask{arr, "", 0, false, false},
		)
	case "kill":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k := range e.Services {
			if len(arg1) > 0 && k == arg1 || len(arg1) == 0 || strings.Contains(arg1, "-") {

				arr := []string{
					"kill",
					k,
				}

				if len(signal) > 0 {
					arr = append(arr, "--signal", signal)
				}

				g.Tasks = append(
					g.Tasks,
					CommandTask{arr, "Killing container " + k + " ...", 0, true, false},
				)
			}
		}
	case "restart":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k := range e.Services {
			if len(arg1) > 0 && k == arg1 || len(arg1) == 0 || strings.Contains(arg1, "-") {

				arr := []string{
					"restart",
					k,
				}

				if len(timeout) > 0 {
					arr = append(arr, "--time", timeout)
				}

				g.Tasks = append(
					g.Tasks,
					CommandTask{arr, "Restarting container " + k + " ...", 0, true, false},
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
					CommandTask{[]string{"stop", "--time", "10", k}, "Stopping container " + k + " ...", 0, true, false},
				)
			}
		}
	case "ps":
		g = Command{
			OutputStatus:   false,
			OutputNewlines: true,
		}

		arr := []string{
			"ps",
		}

		if quiet {
			arr = append(arr, "-q")
		}

		g.Tasks = append(
			g.Tasks,
			CommandTask{arr, "", 0, false, false},
		)
	case "pull":
		g = Command{
			OutputStatus:   false,
			OutputNewlines: true,
		}

		for k := range e.Services {
			if (len(arg1) > 0 && k == arg1) || len(arg1) == 0 {
				g.Tasks = append(
					g.Tasks,
					CommandTask{[]string{"pull", k}, "", 0, true, false},
				)
			}
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
