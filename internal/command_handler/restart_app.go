package command_handler

import (
	"os"
	"os/exec"
	"strings"
)

type RestartApp struct {
	Path    string
	Command string
}

func (r RestartApp) Run() {
	err := os.Chdir(r.Path)
	if err != nil {
		panic(err)
	}

	commandParts := strings.Fields(r.Command)
	if len(commandParts) == 0 {
		panic("No command provided!")
	}

	runApp := exec.Command(commandParts[0], commandParts[1:]...)
	runApp.Stdout = os.Stdout
	runApp.Stderr = os.Stderr

	err = runApp.Run()
}
