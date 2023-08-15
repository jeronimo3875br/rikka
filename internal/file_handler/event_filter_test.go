// event_filter_test.go
package file_handler

import (
	"github.com/fsnotify/fsnotify"
	"testing"
)

func TestShouldIgnoreEvent(t *testing.T) {
	ignoreList := []string{"ignore_me", "temp"}

	tests := []struct {
		eventName string
		expected  bool
	}{
		{"some_file.txt", false},
		{"ignore_me.txt", true},
		{"temp_file.tmp", true},
	}

	for _, test := range tests {
		t.Run(test.eventName, func(t *testing.T) {
			actual := ShouldIgnoreEvent(test.eventName, ignoreList)
			if actual != test.expected {
				t.Errorf("For event name %s, expected %v, but got %v", test.eventName, test.expected, actual)
			}
		})
	}
}

func TestShouldTriggerRestart(t *testing.T) {
	tests := []struct {
		op       fsnotify.Op
		expected bool
	}{
		{fsnotify.Write, true},
		{fsnotify.Create, true},
		{fsnotify.Remove, true},
		{fsnotify.Rename, true},
		{fsnotify.Chmod, false},
	}

	for _, test := range tests {
		t.Run(test.op.String(), func(t *testing.T) {
			actual := ShouldTriggerRestart(test.op)
			if actual != test.expected {
				t.Errorf("For operation %s, expected %v, but got %v", test.op, test.expected, actual)
			}
		})
	}
}
