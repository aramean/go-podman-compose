package main

import (
	"fmt"
	"os"
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
	Newlines      bool
	StatusMessage bool
	Tasks         []CommandTask
}

type CommandTask struct {
	Command []string
}

func main() {
	var arg = os.Args[1:]

	if err := runUsage(arg); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	e := parseYML()

	g := buildCommandNew(e, arg)

	for _, field := range g.Tasks {

		e := executeCommand(field.Command)

		if e.statusCode == 0 {
			if g.StatusMessage {
				fmt.Printf(colorGreen, "[OK] ")
			}
			if g.Newlines {
				fmt.Println(e.message)
			} else {
				fmt.Print(e.message)
			}
		} else if e.statusCode == 1 {
			fmt.Printf(colorRed, e.message+"\n")
		}
	}
}

func buildCommandNew(e map[string]Config, arg []string) Command {

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
			Newlines:      true,
			StatusMessage: true,
		}

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

				g.Tasks = append(
					g.Tasks,
					CommandTask{arr},
				)

			}
		}
	case "down":
		g = Command{
			Newlines:      true,
			StatusMessage: true,
		}

		for _, v := range e {
			for k := range v.Services {
				g.Tasks = append(
					g.Tasks,
					CommandTask{[]string{"stop", "--time", "2", k}},
					CommandTask{[]string{"rm", k}},
				)
			}
		}
	case "restart":
		g = Command{
			Newlines:      true,
			StatusMessage: true,
		}

		for _, v := range e {
			for k := range v.Services {
				if (len(arg1) > 0 && k == arg1) || len(arg1) == 0 {
					g.Tasks = append(
						g.Tasks,
						CommandTask{[]string{"restart", k}},
					)
				}
			}
		}
	case "stop":
		g = Command{
			Newlines:      true,
			StatusMessage: true,
		}

		for _, v := range e {
			for k := range v.Services {
				if (len(arg1) > 0 && k == arg1) || len(arg1) == 0 {
					g.Tasks = append(
						g.Tasks,
						CommandTask{[]string{"stop", "--time", "2", k}},
					)
				}
			}
		}
	case "start":
		g = Command{
			Newlines:      true,
			StatusMessage: true,
		}

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

					g.Tasks = append(
						g.Tasks,
						CommandTask{arr},
					)
				}

			}
		}
	case "ps":
		g = Command{
			StatusMessage: false,
			Tasks: []CommandTask{
				{
					[]string{"ps"},
				},
			},
		}
	case "logs":
		g = Command{
			Newlines:      true,
			StatusMessage: true,
		}

		for _, v := range e {
			for k := range v.Services {
				if (len(arg1) > 0 && k == arg1) || len(arg1) == 0 {
					g.Tasks = append(
						g.Tasks,
						CommandTask{[]string{"logs", k}},
					)
				}
			}
		}
	}

	return g
}
