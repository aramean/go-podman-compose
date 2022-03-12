package main

import (
	"fmt"
)

func newComposeBuildCommand(e *Yaml, l []EnvironmentVariable) Command {

	g := Command{
		OutputStatus:   true,
		OutputNewlines: true,
	}

	for k, v := range e.Services {
		name := getContainerName(v, k)
		if (len(arg1) > 0 && name == arg1) || len(arg1) == 0 {
			arr := []string{
				"build",
				"-t", name,
			}

			build := normalizeValue(v.Build)
			context := "."

			for _, r := range build {
				p := transformPairs(r)
				if p.Key == "dockerfile" {
					arr = append(arr, "-f", p.Value)
				}
				if p.Key == "context" {
					context = p.Value
				}
			}

			if len(build) == 0 {
				arr = append(arr, fmt.Sprint(v.Build))
			}

			arr = append(arr, context)

			g.Tasks = append(
				g.Tasks,
				CommandTask{arr, "Building " + name + " ...", 0, false, false},
			)
		}
	}

	return g
}
