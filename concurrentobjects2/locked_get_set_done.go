package concurrentobjects2

import (
	"sync"
)

// locked get + set/done channelパターン
type LockedGetSetDone struct {
	closing struct {
		ch chan struct{}
		sync.WaitGroup
		sync.Mutex
	}
	prop1 struct {
		s      string
		setch  chan<- string
		donech <-chan struct{}
		sync.RWMutex
	}
}

func NewLockedGetSetDone() *LockedGetSetDone {
	m := &LockedGetSetDone{}
	m.closing.ch = make(chan struct{})

	prop1setch, prop1donech := make(chan string), make(chan struct{})
	m.prop1.setch = prop1setch
	m.prop1.donech = prop1donech

	m.closing.Add(1)
	go func() {
		m.setProc(prop1setch, m.closing.ch, prop1donech)
		m.closing.Done()
	}()

	return m
}

func (m *LockedGetSetDone) Close() {
	m.closing.Lock()
	select {
	case <-m.closing.ch:
		panic("its closed")
	default:
		close(m.closing.ch)
	}
	m.closing.Wait()
	m.closing.Unlock()
}

func (m *LockedGetSetDone) Prop1() (s string) {
	// no lock and using sync/atomic package is also ok
	m.prop1.RLock()
	s = m.prop1.s
	m.prop1.RUnlock()
	return
}

func (m *LockedGetSetDone) SetProp1(val string) {
	select {
	case m.prop1.setch <- val:
		<-m.prop1.donech
	case <-m.closing.ch:
		panic("its closed")
	}
}

func (m *LockedGetSetDone) setProc(in <-chan string, closingIn <-chan struct{}, doneOut chan<- struct{}) {
	if cap(in) != 0 {
		panic("it must be cap 0")
	}
	for {
		select {
		case newProp := <-in:
			// no lock and using sync/atomic package is also ok
			m.prop1.Lock()
			m.prop1.s = newProp
			m.prop1.Unlock()
			doneOut <- struct{}{}
		case <-closingIn:
			return
		}
	}
}
