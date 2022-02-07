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

	for v, k := range envs {
		e.Environment = append(
			e.Environment,
			EnvironmentVariable{Name: k, Value: v},
		)
	}

	return &e
}

func replaceEnvironmentVariable(s string) string {
	split := strings.Split(s, "=")
	return split[0] + "=" + split[1]
}
