package cdkwrapper

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"syscall"
)

// ExecuteCDKCommand executes a CDK CLI command, returns its output and exit status.
func ExecuteCDKCommand(args []string) (string, int, error) {
	var outbuf, errbuf bytes.Buffer
	os.Setenv("FORCE_COLOR", "1")
	cmd := exec.Command("cdk", args...)
	cmd.Stdout = io.MultiWriter(&outbuf, os.Stdout)
	cmd.Stderr = io.MultiWriter(&errbuf, os.Stderr)

	err := cmd.Run()

	output := outbuf.String() + errbuf.String()
	var exitStatus int

	if cmd.ProcessState != nil {
		if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
			exitStatus = status.ExitStatus()
		}
	}

	return output, exitStatus, err
}
