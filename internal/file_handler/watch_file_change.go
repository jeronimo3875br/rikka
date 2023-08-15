package file_handler

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/jeronimo3875br/rikka/internal/command_handler"
	"github.com/jeronimo3875br/rikka/utils"
)

type WatchFiles struct {
	Path    string
	Reflect string
	Ignore  []string
}

// FileChanges monitors file changes and restarts the application when necessary.
func (w *WatchFiles) FileChanges(restartChan <-chan bool) {
	utils.ClearTerminal()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	if err := watcher.Add(w.Path); err != nil {
		panic(err)
	}

	fmt.Printf("\033[32mRikka is running -> \u001B[0m \033[33m%v\u001B[0m\n", w.Path)
	w.RestartApp() // Start the application

	lastRestartTime := time.Now()

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if !ShouldIgnoreEvent(event.Name, w.Ignore) && ShouldTriggerRestart(event.Op) && !ContainsTempPatterns(event.Name) && w.EnoughTimeSinceLastRestart(lastRestartTime) {
				fmt.Printf("\n\033[32mChange detected -> \u001B[0m \033[33m%v\u001B[0m\n\033[32mRestarting application ->\u001B[0m \033[33m%v\u001B[0m\n", event.Name, w.Reflect)

				w.RestartApp() // Restart the application

				lastRestartTime = time.Now()
			}

		case <-restartChan:
			lastRestartTime = time.Now()
		}
	}
}

// EnoughTimeSinceLastRestart checks if enough time has passed since the last restart.
func (w *WatchFiles) EnoughTimeSinceLastRestart(lastRestartTime time.Time) bool {
	elapsed := time.Since(lastRestartTime)
	return elapsed >= 2*time.Second
}

// RestartApp restarts the application.
func (w *WatchFiles) RestartApp() {
	startApp := command_handler.RestartApp{
		Path:    w.Path,
		Command: w.Reflect,
	}
	startApp.Run()
}

// Run starts monitoring file changes.
func (w *WatchFiles) Run() {
	restartChan := make(chan bool)
	go w.FileChanges(restartChan)
	<-make(chan struct{})
}
