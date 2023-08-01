package exec

import (
	"os/exec"
	"strings"
)

// Exec command on operating system
func SoExec(command string) (string, error) {

	args := strings.Split(command, " ")

	saida, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		return string(saida), err
	}

	return string(saida), nil
}
