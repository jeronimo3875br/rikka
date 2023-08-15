package command_handler

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	restorePath, _ := os.Getwd()
	defer func() {
		os.Chdir(restorePath)
	}()

	path := "./"
	command := "echo Hello, World!"

	restartApp := RestartApp{
		path,
		command,
	}
	restartApp.Run()
}

func TestNewRestartApp(t *testing.T) {
	path := "./"
	command := "echo Hello, World!"

	restartApp := RestartApp{
		path,
		command,
	}

	if restartApp.Path != path {
		t.Errorf("Expected Path to be %s, but got %s", path, restartApp.Path)
	}
	if restartApp.Command != command {
		t.Errorf("Expected Command to be %s, but got %s", command, restartApp.Command)
	}
}
