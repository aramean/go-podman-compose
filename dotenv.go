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
	e := EnvironmentVariables{}

	/*
		Load environments from file
	*/
	envsFileEnv, _ := godotenv.Read(fileEnv)
	for k, v := range envsFileEnv {
		e.Environment = append(
			e.Environment,
			EnvironmentVariable{Name: k, Value: v},
		)
	}

	/*
		Load environments from OS
	*/
	envsOs := os.Environ()
	for _, v := range envsOs {
		k := strings.SplitN(v, "=", 2)
		e.Environment = append(
			e.Environment,
			EnvironmentVariable{Name: k[0], Value: k[1]},
		)
	}

	return &e
}
