package observer

import (
	"sync"
	"time"

	"github.com/sdkopen/sdkopen-go/logging"
)

var once sync.Once
var singleInstance *sync.WaitGroup

func GetWaitGroup() *sync.WaitGroup {
	if singleInstance == nil {
		once.Do(func() {
			logging.Debug("Creating single WaitGroup instance now.")
			singleInstance = &sync.WaitGroup{}
		})
	} else {
		logging.Debug("Single WaitGroup instance already created.")
	}

	return singleInstance
}

func WaitRunningTimeout() bool {
	timeout := 10
	c := make(chan struct{})
	go func() {
		defer close(c)
		GetWaitGroup().Wait()
	}()
	select {
	case <-c:
		return false
	case <-time.After(time.Duration(timeout) * time.Second):
		return true // timed out
	}
}
