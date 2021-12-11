package command

import (
	"os/exec"
)

func Execute(statement1 string, args []string) string {

	cmd := exec.Command(statement1, args...)
	stdout, err := cmd.Output()

	if err != nil {
		return err.Error()
	}

	return string(stdout)
}
