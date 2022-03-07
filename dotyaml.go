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
	BlkioConfig     ServicesBlockIO     `yaml:"blkio_config,omitempty"`
	CapAdd          []string            `yaml:"cap_add,omitempty"`
	CapDrop         []string            `yaml:"cap_drop,omitempty"`
	Cpus            float32             `yaml:"cpus,omitempty"`
	Cpuset          string              `yaml:"cpuset,omitempty"`
	CgroupParent    string              `yaml:"cgroup_parent,omitempty"`
	CpuPeriod       string              `yaml:"cpu_period,omitempty"`
	CpuRtRuntime    string              `yaml:"cpu_rt_runtime,omitempty"`
	CpuRtPeriod     string              `yaml:"cpu_rt_period,omitempty"`
	CpuShares       string              `yaml:"cpu_shares,omitempty"`
	CpuQuota        int64               `yaml:"cpu_quota,omitempty"`
	Dns             interface{}         `yaml:"dns,omitempty"`
	DnsSearch       interface{}         `yaml:"dns_search,omitempty"`
	DnsOpt          []string            `yaml:"dns_opt,omitempty"`
	Devices         []string            `yaml:"devices,omitempty"`
	ContainerName   string              `yaml:"container_name,omitempty"`
	Entrypoint      interface{}         `yaml:"entrypoint,omitempty"`
	EnvFile         interface{}         `yaml:"env_file,omitempty"`
	Environment     interface{}         `yaml:"environment,omitempty"`
	Expose          interface{}         `yaml:"expose,omitempty"`
	Healthcheck     ServicesHealthcheck `yaml:"healthcheck,omitempty"`
	Hostname        string              `yaml:"hostname,omitempty"`
	Image           string              `yaml:"image"`
	Init            string              `yaml:"init,omitempty"`
	Ipc             string              `yaml:"ipc,omitempty"`
	Labels          interface{}         `yaml:"labels,omitempty"`
	Logging         ServicesLogging     `yaml:"logging,omitempty"`
	MacAddress      string              `yaml:"mac_address,omitempty"`
	MemLimit        string              `yaml:"mem_limit,omitempty"`
	MemReservation  string              `yaml:"mem_reservation,omitempty"`
	MemSwappiness   int8                `yaml:"mem_swappiness,omitempty"`
	MemSwapLimit    string              `yaml:"memswap_limit,omitempty"`
	OomKill         bool                `yaml:"oom_kill_disable,omitempty"`
	OomScoreAdj     int16               `yaml:"oom_score_adj,omitempty"`
	Platform        string              `yaml:"platform,omitempty"`
	Ports           []string            `yaml:"ports,omitempty"`
	Pid             string              `yaml:"pid,omitempty"`
	PidsLimit       int16               `yaml:"pids_limit,omitempty"`
	Privileged      bool                `yaml:"privileged,omitempty"`
	Readonly        bool                `yaml:"read_only,omitempty"`
	Restart         string              `yaml:"restart,omitempty"`
	Secrets         []ServicesSecrets   `yaml:"secrets"`
	SecurityOpt     []string            `yaml:"security_opt,omitempty"`
	ShmSize         string              `yaml:"shm_size,omitempty"`
	StdinOpen       bool                `yaml:"stdin_open,omitempty"`
	StopGracePeriod string              `yaml:"stop_grace_period,omitempty"`
	StopSignal      string              `yaml:"stop_signal,omitempty"`
	Sysctls         interface{}         `yaml:"sysctls,omitempty"`
	Tmpfs           interface{}         `yaml:"tmpfs,omitempty"`
	Tty             bool                `yaml:"tty,omitempty"`
	Ulimits         ServicesUlimits     `yaml:"ulimits"`
	User            string              `yaml:"user,omitempty"`
	UsernsMode      string              `yaml:"userns_mode,omitempty"`
	Volumes         []string            `yaml:"volumes,omitempty"`
	WorkingDir      string              `yaml:"working_dir,omitempty"`
}

type ServicesLogging struct {
	Driver  string      `yaml:"driver,omitempty"`
	Options interface{} `yaml:"options,omitempty"`
}

type ServicesBlockIO struct {
	Weight          int16                         `yaml:"weight,omitempty"`
	WeightDevice    []ServicesBlockIOWeightDevice `yaml:"weight_device,omitempty"`
	DeviceReadBps   []ServicesBlockIORateDevice   `yaml:"device_read_bps,omitempty"`
	DeviceReadIops  []ServicesBlockIORateDevice   `yaml:"device_read_iops,omitempty"`
	DeviceWriteBps  []ServicesBlockIORateDevice   `yaml:"device_write_bps,omitempty"`
	DeviceWriteIops []ServicesBlockIORateDevice   `yaml:"device_write_iops,omitempty"`
}

type ServicesBlockIOWeightDevice struct {
	Path   string `yaml:"path,omitempty"`
	Weight string `yaml:"weight,omitempty"`
}

type ServicesBlockIORateDevice struct {
	Path string `yaml:"path,omitempty"`
	Rate string `yaml:"rate,omitempty"`
}

type ServicesHealthcheck struct {
	Disable     bool        `yaml:"disable"`
	Test        interface{} `yaml:"test,omitempty"`
	Interval    string      `yaml:"interval,omitempty"`
	Timeout     string      `yaml:"timeout,omitempty"`
	Retries     int8        `yaml:"retries,omitempty"`
	StartPeriod string      `yaml:"start_period,omitempty"`
}

type ServicesSecrets struct {
	Source string `yaml:"source,omitempty"`
	Target string `yaml:"target,omitempty"`
	Uid    string `yaml:"uid,omitempty"`
	Gid    string `yaml:"gid,omitempty"`
	Mode   int16  `yaml:"mode,omitempty"`
}

type ServicesUlimits struct {
	Core         interface{} `yaml:"core,omitempty"`
	Data         interface{} `yaml:"data,omitempty"`
	Fsize        interface{} `yaml:"fsize,omitempty"`
	Memlock      interface{} `yaml:"memlock,omitempty"`
	Nofile       interface{} `yaml:"nofile,omitempty"`
	Rss          interface{} `yaml:"rss,omitempty"`
	Stack        interface{} `yaml:"stack,omitempty"`
	Cpu          interface{} `yaml:"cpu,omitempty"`
	Nproc        interface{} `yaml:"nproc,omitempty"`
	As           interface{} `yaml:"as,omitempty"`
	Maxlogins    interface{} `yaml:"maxlogins,omitempty"`
	Maxsyslogins interface{} `yaml:"maxsyslogins,omitempty"`
	Priority     interface{} `yaml:"priority,omitempty"`
	Locks        interface{} `yaml:"locks,omitempty"`
	Sigpending   interface{} `yaml:"sigpending,omitempty"`
	Msgqueue     interface{} `yaml:"msgqueue,omitempty"`
	Nice         interface{} `yaml:"nice,omitempty"`
	Rtprio       interface{} `yaml:"rtprio,omitempty"`
	Chroot       interface{} `yaml:"chroot,omitempty"`
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
			arr = append(arr, fmt.Sprint(v))
		}
	case map[string]interface{}:
		for k, v := range t {

			var x = ""

			switch v := v.(type) {
			case int:
				x = strconv.FormatInt(int64(v), 10)
			case bool, string:
				x = fmt.Sprint(v)
			}

			arr = append(arr, k+"="+x)
		}
	}

	return arr
}

func normalizeValueColonsPair(field interface{}) string {
	values := normalizeValue(field)
	size := len(values)
	val := ""

	if size == 0 {
		return fmt.Sprint(field)
	}

	for i, r := range values {
		p := transformPairs(r)
		val += p.Value
		if size-1 != i {
			val += ":"
		}
	}
	return val
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
			arr = append(arr, fmt.Sprint(v))
		}
	default:
		arr = append(arr, fmt.Sprint(t))
	}

	return arr
}

func convertToString(p []string) string {
	str := ""
	if len(p) > 0 {
		str = strings.Join(p, " ")
	}
	return str
}

func extractLetters(str string) string {
	re, _ := regexp.Compile(`[^a-zA-Z]+`)
	str = re.ReplaceAllString(str, "")
	return str
}

func extractNumbers(str string) string {
	re, _ := regexp.Compile(`[^\d.-]`)
	str = re.ReplaceAllString(str, "")
	return str
}

func setDurationUnit(str string, unit string) string {
	letter := strings.ToLower(extractLetters(str))
	h, _ := time.ParseDuration(str)

	switch unit {
	case "us":
		if letter != "us" {
			return fmt.Sprint(h.Microseconds())
		}
	case "ms":
		if letter != "ms" {
			return fmt.Sprint(h.Milliseconds())
		}
	case "s":
		if letter != "s" {
			return fmt.Sprint(h.Seconds())
		}
	case "m":
		if letter != "m" {
			return fmt.Sprint(h.Minutes())
		}
	case "h":
		if letter != "h" {
			return fmt.Sprint(h.Hours())
		}
	}

	return extractNumbers(str)
}

func setByteUnit(str string, unit string) string {
	letter := strings.ToLower(extractLetters(str))
	number, _ := strconv.ParseFloat(extractNumbers(str), 64)

	var bytes float64 = number

	switch letter {
	case "k", "kb":
		bytes = number * 1000
	case "m", "mb":
		bytes = number * 1000000
	case "g", "gb":
		bytes = number * 1000000000
	}

	var val = fmt.Sprint(number) + "b"

	switch unit {
	case "k":
		val = fmt.Sprint(bytes/1000) + "k"
	case "m":
		val = fmt.Sprint(bytes/1000000) + "m"
	case "g":
		val = fmt.Sprint(bytes/1000000000) + "g"
	}

	return val
}
