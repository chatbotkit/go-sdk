//go:build windows

package agent

import "os/exec"

// configureCancellation is a no-op on Windows, which has no process-group /
// SIGKILL semantics equivalent to the Unix implementation. Cancellation falls
// back to the default exec.CommandContext behaviour (Process.Kill on the
// top-level process).
func configureCancellation(cmd *exec.Cmd) {}
