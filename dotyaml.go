package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Services struct {
	Dns         interface{} `yaml:"dns,omitempty"`
	Image       string      `yaml:"image"`
	Volumes     []string    `yaml:"volumes,omitempty"`
	Platform    string      `yaml:"platform,omitempty"`
	Ports       []string    `yaml:"ports,omitempty"`
	Restart     string      `yaml:"restart,omitempty"`
	EnvFile     interface{} `yaml:"env_file,omitempty"`
	Environment interface{} `yaml:"environment,omitempty"`
	Init        string      `yaml:"init,omitempty"`
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

func parseYAML(l []EnvironmentVariable) *Config {

	yfile, err2 := ioutil.ReadFile(fileYAML)

	if err2 != nil {
		log.Fatal(err2)
	}

	yfile = replaceEnvironmentVariables(l, yfile)

	var e Config
	err3 := yaml.Unmarshal(yfile, &e)

	if err3 != nil {
		log.Fatal(err3)
	}

	return &e
}

func replaceEnvironmentVariables(l []EnvironmentVariable, yfile []byte) []byte {
	content := string(yfile)

	re := regexp.MustCompile(`\${(.*?)}`)

	replaced := re.ReplaceAllStringFunc(content, func(part string) string {
		plen := len(part)

		if plen > 2 {
			var name = strings.Split(part[2:plen-1], ":")
			for _, k := range l {
				if name[0] == k.Name {
					return k.Value
				}
			}

			if len(name) == 2 {
				return name[1][1:len(name[1])]
			}

		}
		return part
	})

	return []byte(replaced)
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

			var x = ""

			switch v := v.(type) {
			case int:
				x = strconv.FormatInt(int64(v), 10)
			case bool, string:
				x = v.(string)
			}

			arr = append(arr, k+"="+x)
		}
	}

	return arr
}

func convertToArray(t interface{}) []string {
	arr := []string{}

	switch t := t.(type) {
	case []interface{}:
		for _, v := range t {
			arr = append(arr, v.(string))
		}
	default:
		arr = append(arr, t.(string))
	}

	return arr
}
