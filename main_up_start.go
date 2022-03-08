package main

import (
	"fmt"
	"strconv"
)

func newComposeUpStartCommand(e *Yaml, l []EnvironmentVariable) Command {

	g := Command{
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

			if v.StorageOpt.Size != "" {
				arr = append(arr, "--storage-opt", "size="+v.StorageOpt.Size)
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

	return g
}
