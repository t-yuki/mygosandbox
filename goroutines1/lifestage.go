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

// Lifestage is a unidirectional status management helper.
// Created --> Starting --> Started --> Stooped
//         <-> Creating
type Lifestage struct {
	st int32
}

func (s *Lifestage) Stopped() bool {
	// TODO: use CAS
	st := atomic.LoadInt32(&s.st)
	return st == stoppedState
}

func (s *Lifestage) SetCreating() bool {
	return atomic.CompareAndSwapInt32(&s.st, createdState, creatingState)
}

func (s *Lifestage) SetCreated() bool {
	return atomic.CompareAndSwapInt32(&s.st, creatingState, createdState)
}

func (s *Lifestage) SetStarting() bool {
	return atomic.CompareAndSwapInt32(&s.st, createdState, startingState)
}

func (s *Lifestage) SetStarted() bool {
	return atomic.CompareAndSwapInt32(&s.st, startingState, startedState)
}

func (s *Lifestage) SetStopped() bool {
	return atomic.CompareAndSwapInt32(&s.st, startedState, stoppedState)
}
