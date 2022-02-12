package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Services struct {
	Image       string      `yaml:"image"`
	Volumes     []string    `yaml:"volumes"`
	Ports       []string    `yaml:"ports"`
	Restart     string      `yaml:"restart"`
	Environment interface{} `yaml:"environment"`
}

type Volumes struct {
	Mounts []string `yaml:"volumes"`
}

type Ports struct {
	Ports []string `yaml:"ports"`
}

type Restart struct {
	Restart string `yaml:"restart"`
}

type Environment struct {
	Environment interface{} `yaml:"environment"`
}

type Networks struct {
	Networks []string `yaml:"networks"`
}

type Config struct {
	Version  string
	Services map[string]Services
	Volumes  map[string]Volumes
}

func (e *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var services map[string]Services
	var volumes map[string]Volumes

	if err := unmarshal(&services); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}
	if err := unmarshal(&volumes); err != nil {
		if _, ok := err.(*yaml.TypeError); !ok {
			return err
		}
	}

	e.Services = services
	e.Volumes = volumes
	return nil
}

func parseYAML() map[string]Config {
	yfile, err2 := ioutil.ReadFile("docker-compose.yml")

	if err2 != nil {
		log.Fatal(err2)
	}

	//yfile = bytes.Replace(yfile, []byte("mariadb"), []byte("ok"), -1)

	var e map[string]Config
	err3 := yaml.Unmarshal(yfile, &e)

	if err3 != nil {
		log.Fatal(err3)
	}

	return e
}

func convertEnvironmentVariable(t interface{}) []string {
	arr := []string{}

	switch t := t.(type) {
	case []interface{}:
		for _, v := range t {
			arr = append(arr, v.(string))
		}
	case map[string]interface{}:
		for k, v := range t {
			arr = append(arr, k+"="+v.(string))
		}
	}

	return arr
}
