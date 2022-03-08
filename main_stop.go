package main

func newComposeStopCommand(e *Yaml, l []EnvironmentVariable) Command {

	g := Command{
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

	return g
}
