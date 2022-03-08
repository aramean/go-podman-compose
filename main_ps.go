package main

func newComposePsCommand(e *Yaml, l []EnvironmentVariable) Command {

	g := Command{
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

	return g
}
