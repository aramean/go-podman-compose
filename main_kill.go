package main

import (
	"strings"
)

func newComposeKillCommand(e *Yaml, l []EnvironmentVariable) Command {

	g := Command{
		OutputStatus:   true,
		OutputNewlines: true,
	}

	for k, v := range e.Services {
		name := getContainerName(v, k)
		if len(arg1) > 0 && name == arg1 || len(arg1) == 0 || strings.Contains(arg1, "-") {
			arr := []string{
				"kill",
				name,
			}

			if len(signal) > 0 {
				arr = append(arr, "--signal", signal)
			}

			g.Tasks = append(
				g.Tasks,
				CommandTask{arr, "Killing container " + name + " ...", 0, true, false},
			)
		}
	}

	return g
}
