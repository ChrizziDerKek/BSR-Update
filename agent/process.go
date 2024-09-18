package agent

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/shirou/gopsutil/process"
)

func OpenProcess() error {

	var cmd *exec.Cmd

	if *IsServerUpdate {
		if runtime.GOOS == "windows" {
			path, err := filepath.Abs("BSR_Server.exe")
			if err != nil {
				return err
			}

			cmd = exec.Command("cmd", "/c", "start", "cmd", "/k", path)
		} else {
			cmd = exec.Command("mono", "BSR_Server.exe")
		}
	} else {
		if runtime.GOOS == "windows" {
			path, err := filepath.Abs("BSR_Client.exe")
			if err != nil {
				return err
			}

			cmd = exec.Command(path)
		} else {
			cmd = exec.Command("mono", "BSR_Client.exe")
		}
	}

	return cmd.Start()
}

func KillProcess() error {
	if *IsServerUpdate {
		if runtime.GOOS == "windows" {
			return terminateProcess("BSR_Server.exe")
		}
	} else {
		if runtime.GOOS == "windows" {
			return terminateProcess("BSR_Client.exe")
		}
	}

	return terminateProcess("mono")
}

func terminateProcess(name string) error {
	processes, err := process.Processes()
	if err != nil {
		return err
	}
	for _, p := range processes {
		n, err := p.Name()
		if err != nil {
			return err
		}
		if n == name {
			return p.Kill()
		}
	}
	return fmt.Errorf("process not found")
}
