// +build !windows

package oci

import (
	"fmt"
	"os/exec"

	"golang.org/x/sys/unix"
)

func kill(pid int) error {
	err := unix.Kill(pid, unix.SIGKILL)
	if err != nil && err != unix.ESRCH {
		return fmt.Errorf("failed to kill process: %v", err)
	}
	return nil
}

func getExitCode(err error) int32 {
	if exitErr, ok := err.(*exec.ExitError); ok {
		if status, ok := exitErr.Sys().(unix.WaitStatus); ok {
			return int32(status.ExitStatus())
		}
	}
	return -1
}
