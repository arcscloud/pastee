package utl

import "os/exec"

// Run is a small wrapper for exec.Command
func Run(command string, args ...string) ([]byte, error) {
    c := exec.Command(command, args...)
    return c.CombinedOutput()
}
