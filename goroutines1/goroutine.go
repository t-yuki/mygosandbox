package goroutines1

import (
	"sync"
)

type Goroutine struct {
	st Lifestage
	ch chan struct{}
	wg sync.WaitGroup
}

func (s *Goroutine) Start0(f func()) {
	if !s.setStarted() {
		// illegal state
		return
	}
	go func() {
		defer s.wg.Done()
		f()
	}()
}

func (s *Goroutine) Start(f func(<-chan struct{})) {
	if !s.setStarted() {
		// illegal state
		return
	}
	go func() {
		defer s.wg.Done()
		f(s.ch)
	}()
}

func (s *Goroutine) setStarted() bool {
	if !s.st.SetStarting() {
		// illegal state
		return false
	}
	s.ch = make(chan struct{}, 1)
	s.wg.Add(1)
	s.st.SetStarted()
	return true
}

func (s *Goroutine) Stop() {
	if !s.st.SetStopped() {
		// illegal state
		return
	}
	s.ch <- struct{}{}
}

func (s *Goroutine) Wait() {
	if !s.st.Stopped() {
		// illegal state
		return
	}
	s.wg.Wait()
}
