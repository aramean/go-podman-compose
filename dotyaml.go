package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Services struct {
	CapAdd         []string    `yaml:"cap_add,omitempty"`
	CapDrop        []string    `yaml:"cap_drop,omitempty"`
	Cpus           float32     `yaml:"cpus,omitempty"`
	Cpuset         string      `yaml:"cpuset,omitempty"`
	CgroupParent   string      `yaml:"cgroup_parent,omitempty"`
	CpuPeriod      string      `yaml:"cpu_period,omitempty"`
	CpuRtRuntime   string      `yaml:"cpu_rt_runtime,omitempty"`
	CpuRtPeriod    string      `yaml:"cpu_rt_period,omitempty"`
	CpuShares      string      `yaml:"cpu_shares,omitempty"`
	CpuQuota       int64       `yaml:"cpu_quota,omitempty"`
	Dns            interface{} `yaml:"dns,omitempty"`
	DnsSearch      interface{} `yaml:"dns_search,omitempty"`
	DnsOpt         []string    `yaml:"dns_opt,omitempty"`
	Devices        []string    `yaml:"devices,omitempty"`
	ContainerName  string      `yaml:"container_name,omitempty"`
	Entrypoint     interface{} `yaml:"entrypoint,omitempty"`
	EnvFile        interface{} `yaml:"env_file,omitempty"`
	Environment    interface{} `yaml:"environment,omitempty"`
	Expose         interface{} `yaml:"expose,omitempty"`
	Hostname       string      `yaml:"hostname,omitempty"`
	Image          string      `yaml:"image"`
	Init           string      `yaml:"init,omitempty"`
	Ipc            string      `yaml:"ipc,omitempty"`
	Labels         interface{} `yaml:"labels,omitempty"`
	MacAddress     string      `yaml:"mac_address,omitempty"`
	MemLimit       string      `yaml:"mem_limit,omitempty"`
	MemReservation string      `yaml:"mem_reservation,omitempty"`
	MemSwappiness  int8        `yaml:"mem_swappiness,omitempty"`
	MemSwapLimit   string      `yaml:"memswap_limit,omitempty"`
	OomKill        bool        `yaml:"oom_kill_disable,omitempty"`
	OomScoreAdj    int16       `yaml:"oom_score_adj,omitempty"`
	Platform       string      `yaml:"platform,omitempty"`
	Ports          []string    `yaml:"ports,omitempty"`
	Pid            string      `yaml:"pid,omitempty"`
	PidsLimit      int16       `yaml:"pids_limit,omitempty"`
	Privileged     bool        `yaml:"privileged,omitempty"`
	Readonly       bool        `yaml:"read_only,omitempty"`
	Restart        string      `yaml:"restart,omitempty"`
	ShmSize        string      `yaml:"shm_size,omitempty"`
	StopSignal     string      `yaml:"stop_signal,omitempty"`
	Sysctls        interface{} `yaml:"sysctls,omitempty"`
	Tmpfs          interface{} `yaml:"tmpfs,omitempty"`
	Tty            bool        `yaml:"tty,omitempty"`
	User           string      `yaml:"user,omitempty"`
	Volumes        []string    `yaml:"volumes,omitempty"`
	WorkingDir     string      `yaml:"working_dir,omitempty"`
}

type Volumes struct {
	Mounts []string `yaml:"volumes"`
}

type Environment struct {
	Environment interface{} `yaml:"environment"`
}

type Networks struct {
	Networks []string `yaml:"networks"`
}

type Yaml struct {
	Version  string
	Services map[string]Services
	Volumes  map[string]Volumes
}

type YamlPairs struct {
	Key   string
	Value string
}

func parseYAML(l []EnvironmentVariable) *Yaml {

	yfile, err2 := ioutil.ReadFile(fileYAML)

	if err2 != nil {
		log.Fatal(err2)
	}

	yfile = replaceEnvironmentVariables(l, yfile)

	var e Yaml
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

func normalizeValue(t interface{}) []string {
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

func transformPairs(s string) *YamlPairs {
	split := strings.Split(s, "=")
	var k, v string

	if len(split) == 1 {
		k = split[0]
	}

	if len(split) == 2 {
		k = split[0]
		v = split[1]
	}

	p := YamlPairs{Key: k, Value: v}

	return &p
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

func convertToString(p []string) string {
	str := strings.Join(p, " ")
	return str
}

func extractNumbers(str string) string {
	re, _ := regexp.Compile(`[^\d.-]`)
	str = re.ReplaceAllString(str, "")
	return str
}

func setDurationUnit(str string, unit string) string {
	h, _ := time.ParseDuration(str)

	switch unit {
	case "us":
		if strings.Contains(str, "us") {
			return extractNumbers(str)
		}
		return fmt.Sprint(h.Microseconds())
	case "ms":
		if strings.Contains(str, "ms") {
			return extractNumbers(str)
		}
		return fmt.Sprint(h.Milliseconds())
	case "s":
		if strings.Contains(str, "s") {
			return extractNumbers(str)
		}
		return fmt.Sprint(h.Seconds())
	case "m":
		if strings.Contains(str, "m") {
			return extractNumbers(str)
		}
		return fmt.Sprint(h.Minutes())
	case "h":
		if strings.Contains(str, "h") {
			return extractNumbers(str)
		}
		return fmt.Sprint(h.Hours())
	}

	return str
}
