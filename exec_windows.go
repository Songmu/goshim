// +build windows

package goshim

import (
	"os"
	"os/exec"
	"syscall"
)

func init() {
	execFunc = func(argv0 string, argv []string, envv []string) (err error) {
		cmd := exec.Command(argv0, argv...)
		cmd.Env = envv
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		state := cmd.ProcessState
		if state == nil {
			// process not started
			return err
		}
		ws := state.Sys().(syscall.WaitStatus)
		if !ws.Exited() {
			os.Exit(1)
		}
		os.Exit(ws.ExitStatus())
		return nil
	}
}
