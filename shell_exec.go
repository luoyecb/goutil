package goutil

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ShellExecF(format string, args ...interface{}) (string, string, error) {
	return ShellExec(fmt.Sprintf(format, args...))
}

func ShellExec(cmd string) (string, string, error) {
	// fmt.Printf("ShellExec: %s\n", cmd)
	var out bytes.Buffer
	var errOut bytes.Buffer

	command := exec.Command("/bin/bash", "-c", cmd)
	command.Stdout = &out
	command.Stderr = &errOut

	err := command.Run()
	if err != nil {
		return "", errOut.String(), err
	}
	return out.String(), "", nil
}
