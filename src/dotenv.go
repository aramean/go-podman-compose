package main

import (
	"github.com/joho/godotenv"
)

type EnvironmentVariables struct {
	Environment []EnvironmentVariable
}

type EnvironmentVariable struct {
	Name  string
	Value string
}

func loadEnv() *EnvironmentVariables {

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
