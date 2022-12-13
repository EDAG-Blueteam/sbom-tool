package utils

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"runtime"
)

func ExecCmd(name string, args ...string) (string, error) {
	var cmd *exec.Cmd

	if IsWindows() {
		winArgs := append([]string{"/C", name}, args...)
		cmd = exec.Command("cmd", winArgs...)
	} else {
		cmd = exec.Command(name, args...)
	}

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Println(err)
		return "", err
	}

	if err := cmd.Start(); err != nil {
		log.Println(err)
		return "", err
	}

	data, err := io.ReadAll(stdout)

	if err != nil {
		log.Println(err)
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		log.Println(err)
		return "", err
	}

	fmt.Printf("%s\n", string(data))
	return string(data), nil
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}
