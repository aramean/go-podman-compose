package main

func newComposePullCommand(e *Yaml, l []EnvironmentVariable) Command {

	g := Command{
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

	return g
}
