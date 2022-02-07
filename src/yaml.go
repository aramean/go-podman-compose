package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Services struct {
	Image       string   `yaml:"image"`
	Volumes     []string `yaml:"volumes"`
	Ports       []string `yaml:"ports"`
	Environment []string `yaml:"environment"`
}

type Volumes struct {
	Mounts []string `yaml:"volumes"`
}

type Ports struct {
	Ports []string `yaml:"ports"`
}

type Environment struct {
	Environment []string `yaml:"environment"`
}

type Networks struct {
	Networks []string `yaml:"networks"`
}

type Config struct {
	SkipHeaderValidation bool
	Services             map[string]Services
	Volumes              map[string]Volumes
}

func (e *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	/*var params struct {
		SkipHeaderValidation bool `yaml:"skip-header-validation"`
	}
	if err := unmarshal(&params); err != nil {
		return err
	}*/
	var services map[string]Services
	if err := unmarshal(&services); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}
	e.Services = services

	var volumes map[string]Volumes
	//e.SkipHeaderValidation = params.SkipHeaderValidation
	e.Volumes = volumes
	return nil
}

func parseYAML() map[string]Config {
	yfile, err2 := ioutil.ReadFile("docker-compose.yml")

	if err2 != nil {
		log.Fatal(err2)
	}

	var e map[string]Config
	err3 := yaml.Unmarshal(yfile, &e)

	if err3 != nil {
		log.Fatal(err3)
	}

	return e
}
