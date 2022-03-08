package main

// TODO
func newComposeExecCommand(e *Yaml, l []EnvironmentVariable) Command {

	g := Command{
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

	return g
}
