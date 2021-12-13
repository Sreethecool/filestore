package command

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/Sreethecool/filestore/server/utils"
)

//Runs the given command and returns the output
//Handles Piped commands for word count
func Execute(statement string) (string, error) {
	statement = strings.Replace(statement, "| ", "|", -1)
	statement = strings.Replace(statement, " |", "|", -1)
	cmds := strings.Split(statement, "|")
	//fmt.Println(cmds)
	execCommands := []*exec.Cmd{}
	for _, cmd := range cmds {
		cmdArgs := strings.Split(cmd, " ")
		args := []string{}
		if len(cmdArgs) > 1 {
			args = cmdArgs[1:]
		}
		if utils.IsAllowedCommand(cmdArgs[0]) {
			c := exec.Command(cmdArgs[0], args...)
			execCommands = append(execCommands, c)
		}
	}
	count := len(execCommands)
	pipes := make([]*io.PipeWriter, count)
	for i := 0; i < count-1; i++ {
		r, w := io.Pipe()
		execCommands[i].Stdout = w
		execCommands[i+1].Stdin = r
		pipes[i] = w
	}

	var out bytes.Buffer
	var error_buffer bytes.Buffer
	execCommands[count-1].Stdout = &out
	execCommands[count-1].Stderr = &error_buffer
	if err := run(execCommands, pipes); err != nil {
		fmt.Println("error in running pipes", error_buffer.String())
		return "", err
	}
	return out.String(), nil
}

func run(stack []*exec.Cmd, pipes []*io.PipeWriter) (err error) {
	if stack[0].Process == nil {
		if err = stack[0].Start(); err != nil {
			fmt.Println("Error in Starting run")
			return err
		}
	}
	if len(stack) > 1 {
		if err = stack[1].Start(); err != nil {
			fmt.Println("Error in Starting second run ")
			return err
		}
		defer func() {
			if err == nil {
				pipes[0].Close()
				err = run(stack[1:], pipes[1:])
			}
		}()
	}
	return stack[0].Wait()
}
