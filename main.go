package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	debug         = os.Getenv("DEBUG")
	binaryName    = "podman-compose"
	binaryVersion = "1.0.9"
	args          = os.Args[1:]
	detach        bool
	timeout       string
	removeOrphans bool
	signal        string
	tty           bool
	user          string
	quiet         bool
)

const (
	fileYAML         = "docker-compose.yml"
	fileYAMLOverride = "docker-compose.override.yml"
	fileEnv          = ".env"
)

const (
	colorBlue   = "\033[1;34m%s\033[0m"
	colorCyan   = "\033[1;36m%s\033[0m"
	colorGreen  = "\033[1;32m%s\033[0m"
	colorPurple = "\033[1;35m%s\033[0m"
	colorRed    = "\033[1;31m%s\033[0m"
	colorYellow = "\033[1;33m%s\033[0m"
	colorWhite  = "\033[1;37m%s\033[0m"
)

type Command struct {
	OutputStatus   bool
	OutputNewlines bool
	Tasks          []CommandTask
}

func init() {
	if err := runUsage(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	for i, f := range args {
		if f == "-d" || f == "--detach" {
			detach = true
		}
		if f == "-t" || f == "--timeout" {
			timeout = args[(i + 1)]
		}
		if f == "--remove-orphans" {
			removeOrphans = true
		}
		if f == "-q" || f == "--quiet" {
			quiet = true
		}
		if f == "-s" || f == "--signal" {
			signal = args[(i + 1)]
		}
		if f == "-u" || f == "--user" {
			user = args[(i + 1)]
		}
	}
}

func main() {
	l := loadEnvironmentVariables()
	e := parseYAML(l.Environment)
	g := buildCommand(e, l.Environment)

	for _, field := range g.Tasks {

		e := executeCommand(
			field.Command,
			field.OutputCustomMessage,
			field.OutputSingleline,
			field.OutputQuiet,
		)

		output := strings.TrimSpace(e.OutputCustomMessage)

		switch e.OutputStatusCode {
		case 0:
			if !field.OutputQuiet {
				if g.OutputStatus {
					fmt.Printf(colorGreen, "[OK] ")
				}
				if g.OutputNewlines {
					fmt.Println(output)
				} else {
					fmt.Print(output)
				}
			}
		case 1:

			m := strings.Split(output, "\n")

			for i, s := range m {
				if i == 0 {
					if s != "" {
						fmt.Printf(colorRed, "[**] ")
						fmt.Println(s)
					}
				} else {
					fmt.Printf(colorRed, "|**| "+s+"\n")
				}
			}
		}
	}
}

func buildCommand(e *Yaml, l []EnvironmentVariable) Command {

	var arg0, arg1 string

	for i, v := range args {
		if i == 0 {
			arg0 = v
		}
		if i == 1 {
			arg1 = v
		}
	}

	g := Command{}

	switch arg0 {
	case "up", "start":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k, v := range e.Services {
			name := getContainerName(v, k)
			if (len(arg1) > 0 && name == arg1) || len(arg1) == 0 {
				arr := []string{
					"run",
					"--replace",
					"--name", name,
					"--detach",
				}

				if v.BlkioConfig.Weight > 0 {
					arr = append(arr, "--blkio-weight", fmt.Sprint(v.BlkioConfig.Weight))
				}

				for _, r := range v.BlkioConfig.WeightDevice {
					arr = append(arr, "--blkio-weight-device="+r.Path+":"+r.Weight)
				}

				for _, r := range v.BlkioConfig.DeviceWriteBps {
					arr = append(arr, "--device-write-bps="+r.Path+":"+r.Rate)
				}

				for _, r := range v.BlkioConfig.DeviceReadBps {
					arr = append(arr, "--device-read-bps="+r.Path+":"+r.Rate)
				}

				for _, r := range v.BlkioConfig.DeviceWriteIops {
					arr = append(arr, "--device-write-iops="+r.Path+":"+r.Rate)
				}

				for _, r := range v.BlkioConfig.DeviceReadIops {
					arr = append(arr, "--device-read-iops="+r.Path+":"+r.Rate)
				}

				for i := range v.CapAdd {
					arr = append(arr, "--cap-add", v.CapAdd[i])
				}

				for i := range v.CapDrop {
					arr = append(arr, "--cap-drop", v.CapDrop[i])
				}

				if v.Cpus > 0 {
					arr = append(arr, "--cpus", fmt.Sprint(v.Cpus))
				}

				if v.Cpuset != "" {
					arr = append(arr, "--cpuset-cpus", fmt.Sprint(v.Cpuset))
				}

				if v.CgroupParent != "" {
					arr = append(arr, "--cgroup-parent", v.CgroupParent)
				}

				for _, r := range normalizeValue(v.Labels) {
					r := transformPairs(r)
					arr = append(arr, "--label", r.Key+"="+r.Value)
				}

				for i := range v.Ports {
					arr = append(arr, "--publish", v.Ports[i])
				}

				for i := range v.Volumes {
					arr = append(arr, "--volume", v.Volumes[i])
				}

				if v.CpuPeriod != "" {
					arr = append(arr, "--cpu-period", setDurationUnit(v.CpuPeriod, "ms"))
				}

				if v.CpuRtPeriod != "" {
					arr = append(arr, "--cpu-rt-period", setDurationUnit(v.CpuRtPeriod, "ms"))
				}

				if v.CpuRtRuntime != "" {
					arr = append(arr, "--cpu-rt-runtime", setDurationUnit(v.CpuRtRuntime, "ms"))
				}

				if v.CpuShares != "" {
					arr = append(arr, "--cpu-shares", v.CpuShares)
				}

				if v.CpuQuota > 0 {
					arr = append(arr, "--cpu-quota", fmt.Sprint(v.CpuQuota))
				}

				for i := range v.Devices {
					arr = append(arr, "--device", v.Devices[i])
				}

				if v.Dns != nil {
					for _, r := range convertToArray(v.Dns) {
						arr = append(arr, "--dns", r)
					}
				}

				if v.DnsSearch != nil {
					for _, r := range convertToArray(v.DnsSearch) {
						arr = append(arr, "--dns-search", r)
					}
				}

				for i := range v.DnsOpt {
					arr = append(arr, "--dns-opt", v.DnsOpt[i])
				}

				if v.Entrypoint != nil {
					p := normalizeValue(v.Entrypoint)
					arr = append(arr, "--entrypoint", "'"+convertToString(p)+"'")
				}

				for _, r := range normalizeValue(v.Environment) {
					p := transformPairs(r)
					arr = append(arr, "--env", p.Key+"="+p.Value)
				}

				if v.EnvFile != nil {
					for _, r := range convertToArray(v.EnvFile) {
						arr = append(arr, "--env-file", r)
					}
				}

				if v.Expose != nil {
					for _, r := range convertToArray(v.Expose) {
						arr = append(arr, "--expose", r)
					}
				}

				if v.Healthcheck.Test != nil {
					p := normalizeValue(v.Healthcheck.Test)
					arr = append(arr, "--health-cmd", "'"+convertToString(p)+"'")
				}

				if v.Healthcheck.Interval != "" {
					arr = append(arr, "--health-interval", v.Healthcheck.Interval)
				}

				if v.Healthcheck.Timeout != "" {
					arr = append(arr, "--health-timeout", v.Healthcheck.Timeout)
				}

				if v.Healthcheck.StartPeriod != "" {
					arr = append(arr, "--health-start-period", v.Healthcheck.StartPeriod)
				}

				if v.Healthcheck.Disable {
					arr = append(arr, "--no-healthcheck")
				}

				if v.Hostname != "" {
					arr = append(arr, "--hostname", v.Hostname)
				}

				if v.Init != "" {
					if _, err := strconv.ParseBool(v.Init); err == nil {
						arr = append(arr, "--init")
					} else {
						arr = append(arr, "--init-path", v.Init)
					}
				}

				if v.Ipc != "" {
					arr = append(arr, "--ipc", v.Ipc)
				}

				if v.Logging.Driver != "" {
					arr = append(arr, "--log-driver", v.Logging.Driver)
				}

				if v.Logging.Options != nil {
					for _, r := range normalizeValue(v.Logging.Options) {
						p := transformPairs(r)
						arr = append(arr, "--log-opt", p.Key+"="+p.Value)
					}
				}

				if v.MacAddress != "" {
					arr = append(arr, "--mac-address", v.MacAddress)
				}

				if v.MemLimit != "" {
					arr = append(arr, "--memory", v.MemLimit)
				}

				if v.MemSwappiness > 0 || v.MemSwappiness == -1 {
					arr = append(arr, "--memory-swappiness", fmt.Sprint(v.MemSwappiness))
				}

				if v.MemSwapLimit != "" {
					arr = append(arr, "--memory-swap", v.MemSwapLimit)
				}

				if v.MemReservation != "" {
					arr = append(arr, "--memory-reservation", v.MemReservation)
				}

				if v.OomKill {
					arr = append(arr, "--oom-kill-disable")
				}

				if v.OomScoreAdj <= -1000 || v.OomScoreAdj >= 1000 {
					arr = append(arr, "--oom-score-adj", fmt.Sprint(v.OomScoreAdj))
				}

				if v.Pid != "" {
					arr = append(arr, "--pid", v.Pid)
				}

				if v.PidsLimit < -2 || v.PidsLimit != 0 {
					arr = append(arr, "--pids-limit", fmt.Sprint(v.PidsLimit))
				}

				if v.Platform != "" {
					arr = append(arr, "--platform", v.Platform)
				}

				if v.Privileged {
					arr = append(arr, "--privileged")
				}

				if v.Readonly {
					arr = append(arr, "--read-only")
				}

				if v.Restart != "" {
					arr = append(arr, "--restart", v.Restart)
				}

				if v.SecurityOpt != nil {
					for _, r := range v.SecurityOpt {
						arr = append(arr, "--security-opt", r)
					}
				}

				if v.ShmSize != "" {
					arr = append(arr, "--shm-size", v.ShmSize)
				}

				if v.StdinOpen {
					arr = append(arr, "--interactive")
				}

				if v.StopSignal != "" {
					arr = append(arr, "--stop-signal", v.StopSignal)
				}

				for _, r := range normalizeValue(v.Sysctls) {
					p := transformPairs(r)
					arr = append(arr, "--sysctl", p.Key+"="+p.Value)
				}

				if v.Tmpfs != nil {
					for _, r := range convertToArray(v.Tmpfs) {
						arr = append(arr, "--tmpfs", r)
					}
				}

				if v.Tty {
					arr = append(arr, "--tty")
				}

				if v.Ulimits.Core != nil {
					values := normalizeValueColonsPair(v.Ulimits.Core)
					arr = append(arr, "--ulimit", "core="+values)
				}

				if v.Ulimits.Data != nil {
					values := normalizeValueColonsPair(v.Ulimits.Data)
					arr = append(arr, "--ulimit", "data="+values)
				}

				if v.Ulimits.Fsize != nil {
					values := normalizeValueColonsPair(v.Ulimits.Fsize)
					arr = append(arr, "--ulimit", "fsize="+values)
				}

				if v.Ulimits.Memlock != nil {
					values := normalizeValueColonsPair(v.Ulimits.Memlock)
					arr = append(arr, "--ulimit", "memlock="+values)
				}

				if v.Ulimits.Nofile != nil {
					values := normalizeValueColonsPair(v.Ulimits.Nofile)
					arr = append(arr, "--ulimit", "nofile="+values)
				}

				if v.Ulimits.Rss != nil {
					values := normalizeValueColonsPair(v.Ulimits.Rss)
					arr = append(arr, "--ulimit", "rss="+values)
				}

				if v.Ulimits.Stack != nil {
					values := normalizeValueColonsPair(v.Ulimits.Stack)
					arr = append(arr, "--ulimit", "stack="+values)
				}

				if v.Ulimits.Cpu != nil {
					values := normalizeValueColonsPair(v.Ulimits.Cpu)
					arr = append(arr, "--ulimit", "cpu="+values)
				}

				if v.Ulimits.Nproc != nil {
					values := normalizeValueColonsPair(v.Ulimits.Nproc)
					arr = append(arr, "--ulimit", "nproc="+values)
				}

				if v.Ulimits.As != nil {
					values := normalizeValueColonsPair(v.Ulimits.As)
					arr = append(arr, "--ulimit", "as="+values)
				}

				if v.Ulimits.Maxlogins != nil {
					values := normalizeValueColonsPair(v.Ulimits.Maxlogins)
					arr = append(arr, "--ulimit", "maxlogins="+values)
				}

				if v.Ulimits.Maxsyslogins != nil {
					values := normalizeValueColonsPair(v.Ulimits.Maxsyslogins)
					arr = append(arr, "--ulimit", "maxsyslogins="+values)
				}

				if v.Ulimits.Priority != nil {
					values := normalizeValueColonsPair(v.Ulimits.Priority)
					arr = append(arr, "--ulimit", "priority="+values)
				}

				if v.Ulimits.Locks != nil {
					values := normalizeValueColonsPair(v.Ulimits.Locks)
					arr = append(arr, "--ulimit", "locks="+values)
				}

				if v.Ulimits.Sigpending != nil {
					values := normalizeValueColonsPair(v.Ulimits.Sigpending)
					arr = append(arr, "--ulimit", "sigpending="+values)
				}

				if v.Ulimits.Msgqueue != nil {
					values := normalizeValueColonsPair(v.Ulimits.Msgqueue)
					arr = append(arr, "--ulimit", "msgqueue="+values)
				}

				if v.Ulimits.Nice != nil {
					values := normalizeValueColonsPair(v.Ulimits.Nice)
					arr = append(arr, "--ulimit", "nice="+values)
				}

				if v.Ulimits.Rtprio != nil {
					values := normalizeValueColonsPair(v.Ulimits.Rtprio)
					arr = append(arr, "--ulimit", "rtprio="+values)
				}

				if v.Ulimits.Chroot != nil {
					values := normalizeValueColonsPair(v.Ulimits.Chroot)
					arr = append(arr, "--ulimit", "chroot="+values)
				}

				if v.User != "" {
					arr = append(arr, "--user", v.User)
				}

				if v.UsernsMode != "" {
					arr = append(arr, "--userns", v.UsernsMode)
				}

				if v.WorkingDir != "" {
					arr = append(arr, "--workdir", v.WorkingDir)
				}

				arr = append(arr, v.Image)

				g.Tasks = append(
					g.Tasks,
					CommandTask{arr, "Starting container " + name + " ...", 0, false, false},
				)
			}
		}

		/*if !detach {

			g = Command{
				OutputStatus:   false,
				OutputNewlines: true,
			}

			for k := range e.Services {
				if (len(arg1) > 0 && k == arg1) || len(arg1) == 0 {
					g.Tasks = append(
						g.Tasks,
						CommandTask{[]string{"logs", k}, "", 0, true, false},
					)
				}
			}

		}*/

		for k := range e.Volumes {
			g.Tasks = append(
				g.Tasks,
				CommandTask{[]string{"volume", "create", k}, "", 0, true, true},
			)
		}
	case "down":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k, v := range e.Services {
			name := getContainerName(v, k)
			g.Tasks = append(
				g.Tasks,
				CommandTask{[]string{"stop", "--time", "2", name}, "Stopping container " + name + " ...", 0, true, false},
				CommandTask{[]string{"rm", name}, "", 0, true, true},
			)
		}
	case "exec":
		g = Command{
			OutputStatus:   false,
			OutputNewlines: true,
		}

		exec := ""
		service := ""

		for i, f := range args {
			for k := range e.Services {
				if f == k {
					service = args[i]
					exec = args[(i + 1)]

					fmt.Print("hej")
				}
			}
		}

		arr := []string{
			"exec",
			"-it",
		}

		/*if detach {
			arr = append(arr, "--detach")
		}

		if tty {
			arr = append(arr, "--tty")
		}

		if len(user) > 0 {
			arr = append(arr, "--user", "root")
		}
		*/
		arr = append(arr, service, exec)

		g.Tasks = append(
			g.Tasks,
			CommandTask{arr, "", 0, false, false},
		)
	case "kill":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k, v := range e.Services {
			name := getContainerName(v, k)
			if len(arg1) > 0 && name == arg1 || len(arg1) == 0 || strings.Contains(arg1, "-") {
				arr := []string{
					"kill",
					name,
				}

				if len(signal) > 0 {
					arr = append(arr, "--signal", signal)
				}

				g.Tasks = append(
					g.Tasks,
					CommandTask{arr, "Killing container " + name + " ...", 0, true, false},
				)
			}
		}
	case "restart":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k, v := range e.Services {
			name := getContainerName(v, k)
			if len(arg1) > 0 && name == arg1 || len(arg1) == 0 || strings.Contains(arg1, "-") {
				arr := []string{
					"restart",
					name,
				}

				if len(timeout) > 0 {
					arr = append(arr, "--time", timeout)
				}

				g.Tasks = append(
					g.Tasks,
					CommandTask{arr, "Restarting container " + name + " ...", 0, true, false},
				)
			}
		}
	case "stop":
		g = Command{
			OutputStatus:   true,
			OutputNewlines: true,
		}

		for k, v := range e.Services {
			name := getContainerName(v, k)
			if (len(arg1) > 0 && name == arg1) || len(arg1) == 0 {

				arr := []string{
					"stop",
					name,
				}

				t := "10"
				if timeout != "" {
					t = timeout
				} else if v.StopGracePeriod != "" {
					t = setDurationUnit(v.StopGracePeriod, "s")
				}

				arr = append(arr, "-t", t)

				g.Tasks = append(
					g.Tasks,
					CommandTask{arr, "Stopping container " + name + " ...", 0, true, false},
				)
			}
		}
	case "ps":
		g = Command{
			OutputStatus:   false,
			OutputNewlines: true,
		}

		arr := []string{
			"ps",
		}

		if quiet {
			arr = append(arr, "-q")
		}

		g.Tasks = append(
			g.Tasks,
			CommandTask{arr, "", 0, false, false},
		)
	case "pull":
		g = Command{
			OutputStatus:   false,
			OutputNewlines: true,
		}

		for k, v := range e.Services {
			name := getContainerName(v, k)
			if (len(arg1) > 0 && name == arg1) || len(arg1) == 0 {
				g.Tasks = append(
					g.Tasks,
					CommandTask{[]string{"pull", name}, "", 0, true, false},
				)
			}
		}
	case "logs":
		g = Command{
			OutputStatus:   false,
			OutputNewlines: true,
		}

		for k, v := range e.Services {
			name := getContainerName(v, k)
			if (len(arg1) > 0 && k == arg1) || len(arg1) == 0 {
				g.Tasks = append(
					g.Tasks,
					CommandTask{[]string{"logs", name}, "", 0, true, false},
				)
			}
		}
	}

	return g
}

func getContainerName(v Services, name string) string {
	if v.ContainerName != "" {
		name = v.ContainerName
	}
	return name
}
