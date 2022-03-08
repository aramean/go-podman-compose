package main

import "strings"

func newComposeRestartCommand(e *Yaml, l []EnvironmentVariable) Command {

	g := Command{
		OutputStatus:   true,
		OutputNewlines: true,
	}

	for k, v := range e.Services {
		name := getContainerName(v, k)
		if len(arg1) > 0 && name == arg1 || len(arg1) == 0 || strings.Contains(arg1, "-") {
			arr := []string{
				"restart",
				name,
			}

			if len(timeout) > 0 {
				arr = append(arr, "--time", timeout)
			}

			g.Tasks = append(
				g.Tasks,
				CommandTask{arr, "Restarting container " + name + " ...", 0, true, false},
			)
		}
	}

	return g
}
