package cmdutils

import "os/exec"

// GetCmdInstance function to create a CMD instance depending on the OS platform
func GetCmdInstance(command, platform string) *exec.Cmd {
	var shell string
	var flagForShell string

	switch platform {
	case "windows":
		shell = "cmd"
		flagForShell = "/c"
	case "linux":
		shell = "bash"
		flagForShell = "-c"
	case "darwin":
		shell = "bash"
		flagForShell = "-c"
	}

	return exec.Command(shell, flagForShell, command)
}
