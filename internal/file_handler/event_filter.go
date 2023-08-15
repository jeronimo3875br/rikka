package file_handler

import (
	"github.com/fsnotify/fsnotify"
	"strings"
)

func ShouldIgnoreEvent(eventName string, ignoreList []string) bool {
	for _, ignore := range ignoreList {
		if strings.Contains(eventName, ignore) {
			return true
		}
	}
	return false
}

func ShouldTriggerRestart(eventOp fsnotify.Op) bool {
	return eventOp&fsnotify.Write == fsnotify.Write ||
		eventOp&fsnotify.Create == fsnotify.Create ||
		eventOp&fsnotify.Remove == fsnotify.Remove ||
		eventOp&fsnotify.Rename == fsnotify.Rename
}

func ContainsTempPatterns(filename string) bool {
	tempPatterns := []string{"~", ".tmp", ".temp"}
	for _, pattern := range tempPatterns {
		if strings.Contains(filename, pattern) {
			return true
		}
	}
	return false
}
