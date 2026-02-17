package observer

import (
	"sync"
	"testing"
)

type mockObserver struct {
	closed bool
}

func (m *mockObserver) Close() {
	m.closed = true
}

func newTestService() *service {
	return &service{
		observers: make([]Observer, 0),
		mx:        &sync.Mutex{},
	}
}

func TestAttach_AddsObserver(t *testing.T) {
	svc := newTestService()
	obs := &mockObserver{}

	err := svc.attach(obs)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(svc.observers) != 1 {
		t.Fatalf("expected 1 observer, got %d", len(svc.observers))
	}
}

func TestAttach_MultipleObservers(t *testing.T) {
	svc := newTestService()

	for i := 0; i < 5; i++ {
		err := svc.attach(&mockObserver{})
		if err != nil {
			t.Fatalf("expected no error on attach %d, got %v", i, err)
		}
	}
	if len(svc.observers) != 5 {
		t.Fatalf("expected 5 observers, got %d", len(svc.observers))
	}
}

func TestAttach_RejectsAfterShutdown(t *testing.T) {
	svc := newTestService()
	svc.isShuttingDown = true

	err := svc.attach(&mockObserver{})
	if err == nil {
		t.Fatal("expected error when attaching after shutdown, got nil")
	}
}

func TestNotify_ClosesAllObservers(t *testing.T) {
	svc := newTestService()
	obs1 := &mockObserver{}
	obs2 := &mockObserver{}
	obs3 := &mockObserver{}

	svc.attach(obs1)
	svc.attach(obs2)
	svc.attach(obs3)

	svc.notify()

	if !obs1.closed {
		t.Fatal("expected observer 1 to be closed")
	}
	if !obs2.closed {
		t.Fatal("expected observer 2 to be closed")
	}
	if !obs3.closed {
		t.Fatal("expected observer 3 to be closed")
	}
}

func TestNotify_SetsShuttingDown(t *testing.T) {
	svc := newTestService()
	svc.notify()

	if !svc.isShuttingDown {
		t.Fatal("expected isShuttingDown to be true after notify")
	}
}

func TestNotify_ThenAttach_ReturnsError(t *testing.T) {
	svc := newTestService()
	svc.notify()

	err := svc.attach(&mockObserver{})
	if err == nil {
		t.Fatal("expected error when attaching after notify, got nil")
	}
}

func TestNotify_NoObservers_DoesNotPanic(t *testing.T) {
	svc := newTestService()
	svc.notify()
}

func TestAttach_ConcurrentSafety(t *testing.T) {
	svc := newTestService()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			svc.attach(&mockObserver{})
		}()
	}
	wg.Wait()

	if len(svc.observers) != 100 {
		t.Fatalf("expected 100 observers, got %d", len(svc.observers))
	}
}
