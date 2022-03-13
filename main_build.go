package main

import (
	"fmt"
	"strings"
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
				if p.Key == "args" {
					a := strings.Split(p.Value, delimiter)
					for _, r := range a {
						if len(r) > 0 {
							p := transformPairs(r)
							arr = append(arr, "--build-arg", p.Key+"="+p.Value)
						}
					}
				}
				if p.Key == "labels" {
					a := strings.Split(p.Value, delimiter)
					for _, r := range a {
						if len(r) > 0 {
							p := transformPairs(r)
							arr = append(arr, "--label", p.Key+"="+p.Value)
						}
					}
				}
				if p.Key == "shm_size" {
					arr = append(arr, "--shm-size", p.Value)
				}
				if p.Key == "target" {
					arr = append(arr, "--target", p.Value)
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
