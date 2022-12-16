package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func ExecCmd(name string, args ...string) (string, error) {
	var cmd *exec.Cmd

	if IsWindows() {
		// mvn what ever
		// cmd.exe /C mvn what ever
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

/*
Create a working directory folder if not exist

@param path folder to be created
@return nil
*/
func CreateFolder(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
