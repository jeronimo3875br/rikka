package file_handler

import (
	"testing"
	"time"
)

func TestEnoughTimeSinceLastRestart(t *testing.T) {
	tests := []struct {
		lastRestartTime time.Time
		expected        bool
	}{
		{time.Now().Add(-3 * time.Second), true},
		{time.Now().Add(-1 * time.Second), false},
	}

	watchFiles := WatchFiles{}

	for _, test := range tests {
		t.Run(test.lastRestartTime.String(), func(t *testing.T) {
			actual := watchFiles.enoughTimeSinceLastRestart(test.lastRestartTime)
			if actual != test.expected {
				t.Errorf("For lastRestartTime %v, expected %v, but got %v", test.lastRestartTime, test.expected, actual)
			}
		})
	}
}
