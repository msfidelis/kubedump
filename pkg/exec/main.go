package exec

import (
	"os/exec"
	"strings"
)

// SoExec command on operating system
func SoExec(command string) (string, error) {
	args := strings.Split(command, " ")
	output, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		return string(output), err
	}
	return string(output), nil
}
