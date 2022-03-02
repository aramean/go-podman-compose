package main

import (
	"os"
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
	envsOs := os.Environ()
	envsFileEnv, _ := godotenv.Read(fileEnv)

	e := EnvironmentVariables{}

	for _, v := range envsOs {
		k := strings.SplitN(v, "=", 2)
		e.Environment = append(
			e.Environment,
			EnvironmentVariable{Name: k[0], Value: k[1]},
		)
	}

	for k, v := range envsFileEnv {
		e.Environment = append(
			e.Environment,
			EnvironmentVariable{Name: k, Value: v},
		)
	}

	return &e
}

func transformEnvironmentVariable(s string, l []EnvironmentVariable) *EnvironmentVariable {
	split := strings.Split(s, "=")
	k := split[0]
	v := split[1]

	e := EnvironmentVariable{Name: k, Value: v}

	return &e
}
