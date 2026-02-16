package observer

import (
	"sync"
	"testing"
)

func TestGetWaitGroup_ReturnsSameInstance(t *testing.T) {
	// Reset singleton for this test
	once = sync.Once{}
	singleInstance = nil

	wg1 := GetWaitGroup()
	wg2 := GetWaitGroup()

	if wg1 != wg2 {
		t.Fatal("expected GetWaitGroup to return the same instance")
	}
}

func TestGetWaitGroup_NotNil(t *testing.T) {
	once = sync.Once{}
	singleInstance = nil

	wg := GetWaitGroup()
	if wg == nil {
		t.Fatal("expected GetWaitGroup to return non-nil WaitGroup")
	}
}

func TestWaitRunningTimeout_NoWork_ReturnsFalse(t *testing.T) {
	once = sync.Once{}
	singleInstance = nil

	result := WaitRunningTimeout()
	if result {
		t.Fatal("expected WaitRunningTimeout to return false when no work is pending")
	}
}
