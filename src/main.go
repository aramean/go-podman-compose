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

func main() {
	var arg = os.Args[1:]

	if err := root(arg); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	e := parseYML()
	f := buildCommand(e, arg)

	for _, v := range f {

		e := executeCommand(v)

		if e.status == 0 {
			fmt.Printf(colorGreen, "[OK] ")
			fmt.Println(e.message)
		} else if e.status == 1 {
			fmt.Printf(colorRed, e.message+"\n")
		}
	}
}

func buildCommand(e map[string]Config, arg []string) [][]string {

	var arg0, arg1 = "", ""

	for i, v := range arg {
		if i == 0 {
			arg0 = v
		}
		if i == 1 {
			arg1 = v
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
