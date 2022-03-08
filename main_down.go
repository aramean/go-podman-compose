package main

func newComposeDownCommand(e *Yaml, l []EnvironmentVariable) Command {

	g := Command{
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

	return g
}
