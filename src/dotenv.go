package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func loadEnv() {

	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	name := envs["NAME"]
	editor := envs["EDITOR"]

	fmt.Printf("%s uses %s\n", name, editor)
}
