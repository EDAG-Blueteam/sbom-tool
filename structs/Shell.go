package structs

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sbom-tool/console"
	"strings"
)

type Shell struct {
	Folder   string   `json:"folder"`
	Warnings []string `json:"warnings"`
}

func NewShell(folder string) Shell {

	var shell Shell

	cwd, err1 := os.Getwd()

	if err1 == nil {

		stat, err2 := os.Stat(filepath.Join(cwd, folder))

		if err2 == nil && stat.IsDir() {
			shell.Folder = filepath.Join(cwd, folder)
		} else {
			console.Warn("Shell: \"" + folder + "\" does not exist!")
		}

	}

	return shell

}

func (shell *Shell) Execute(command string, arguments []string) (string, error) {

	var stdout string = ""
	var err error = nil

	if shell.Folder != "" {

		if runtime.GOOS == "windows" {

			// TODO: Support Windows' %PATH% environment variable

			arguments = append([]string{"/C", command}, arguments...)
			command = "C:\\Windows\\system32\\cmd.exe"
		}

		console.Warn(command)
		console.Warn(strings.Join(arguments, ","))
		console.Warn(shell.Folder)

		cmd := exec.Command(command, arguments...)
		cmd.Dir = shell.Folder
		cmd_out, cmd_err := cmd.Output()

		stdout = string(cmd_out)
		err = cmd_err

		if err != nil {
			console.Error(cmd_err.Error())
		}

	} else {
		console.Warn("Shell: Forgot to set folder!")
	}

	return stdout, err

}
