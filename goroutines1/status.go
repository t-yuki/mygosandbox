package goroutines1

import (
	"sync/atomic"
)

const (
	createdState = iota
	creatingState
	startingState
	startedState
	stoppedState
)

// Status is a unidirectional status management helper.
// Created --> Starting --> Started --> Stooped
//         <-> Creating
type Status struct {
	st int32
}

func (s *Status) Stopped() bool {
	// TODO: use CAS
	st := atomic.LoadInt32(&s.st)
	return st == stoppedState
}

func (s *Status) SetCreating() bool {
	return atomic.CompareAndSwapInt32(&s.st, createdState, creatingState)
}

func (s *Status) SetCreated() bool {
	return atomic.CompareAndSwapInt32(&s.st, creatingState, createdState)
}

func (s *Status) SetStarting() bool {
	return atomic.CompareAndSwapInt32(&s.st, createdState, startingState)
}

func (s *Status) SetStarted() bool {
	return atomic.CompareAndSwapInt32(&s.st, startingState, startedState)
}

func (s *Status) SetStopped() bool {
	return atomic.CompareAndSwapInt32(&s.st, startedState, stoppedState)
}
