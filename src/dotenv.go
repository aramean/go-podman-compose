package main

import (
	"strings"

	"github.com/joho/godotenv"
)

type EnvironmentVariables struct {
	Environment []EnvironmentVariable
}

type EnvironmentVariable struct {
	Name  string
	Value string
}

func loadEnvironmentVariables() *EnvironmentVariables {

	envs, _ := godotenv.Read(".env")

	e := EnvironmentVariables{}

	for k, v := range envs {
		e.Environment = append(
			e.Environment,
			EnvironmentVariable{Name: k, Value: v},
		)
	}

	return &e
}

func replaceEnvironmentVariable(s string, l []EnvironmentVariable) *EnvironmentVariable {
	split := strings.Split(s, "=")
	k := split[0]
	v := split[1]

	vfirst := v[0:1]
	vlast := v[len(v)-1:]

	if vfirst == "{" && vlast == "}" {
		for _, i := range l {
			if k == i.Name {
				if k == i.Name {
					v = i.Value
				}
			}
		}
	}

	e := EnvironmentVariable{Name: k, Value: v}

	return &e
}
