package concurrentobjects2

import (
	"sync"
	"time"
)

// single subscriber パターン
type SingleSubscriber struct {
	closing struct {
		ch chan struct{}
		sync.WaitGroup
		sync.Mutex
	}
	ev1 struct {
		subscribech chan<- chan<- string
	}
}

func NewSingleSubscriber() *SingleSubscriber {
	m := &SingleSubscriber{}
	m.closing.ch = make(chan struct{})

	ev1subscribech := make(chan chan<- string)
	m.ev1.subscribech = ev1subscribech

	m.closing.Add(1)
	go func() {
		m.ev1Proc(ev1subscribech, m.closing.ch)
		m.closing.Done()
	}()

	return m
}

func (m *SingleSubscriber) Close() {
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

func (m *SingleSubscriber) Subscribe(val chan<- string) {
	// no check for double subscribe
	select {
	case m.ev1.subscribech <- val:
		break
	case <-m.closing.ch:
		panic("its closed")
	}
}

func (m *SingleSubscriber) Unsubscribe() {
	// no check for double unsubscribe
	select {
	case m.ev1.subscribech <- nil:
		break
	case <-m.closing.ch:
		panic("its closed")
	}
}

func (*SingleSubscriber) ev1Proc(in <-chan chan<- string, closingIn <-chan struct{}) {
	var sub chan<- string
	for {
		select {
		case <-time.After(time.Second):
			// new event
			if sub == nil {
				continue
			}
			select {
			case <-closingIn:
				return
			case sub <- "ping":
				continue
			}
		case newsub := <-in:
			if newsub == sub {
				// skip or not?
			}
			sub = newsub
			if sub == nil {
				continue
			}
			// new sub
			select {
			case <-closingIn:
				return
			case sub <- "initial":
				continue
			}
		case <-closingIn:
			return
		}
	}
}
