//go:build !windows

package agent

import (
	"os/exec"
	"syscall"
)

// configureCancellation starts the command in its own process group and, on
// cancellation, sends SIGKILL to the entire group (negative PID) so that all
// child processes are terminated - not just the top-level shell.
func configureCancellation(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Cancel = func() error {
		return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	}
}
