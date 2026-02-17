package observer

import (
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sdkopen/sdkopen-go/logging"
)

type subject interface {
	attach(observer Observer) error
	notify()
}

var services subject

func Initialize() {
	ch := make(chan os.Signal, 1)
	services = &service{
		observers: make([]Observer, 0, 0),
		mx:        &sync.Mutex{},
	}
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL, os.Interrupt)

	go func() {
		sig := <-ch
		logging.Warn("notify shutdown: %+v", sig)
		services.notify()
	}()
}

func Attach(o Observer) error {
	return services.attach(o)
}

type service struct {
	observers      []Observer
	isShuttingDown bool
	mx             *sync.Mutex
}

func (s *service) attach(observer Observer) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	if s.isShuttingDown {
		logging.Warn("Ignoring new observer after shutdown signal: %T", observer)
		return errors.New("ignoring new observer after shutdown signal")
	}

	s.observers = append(s.observers, observer)
	return nil
}

func (s *service) notify() {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.isShuttingDown = true
	for _, observer := range s.observers {
		observer.Close()
	}
}
