package concurrentobjects2

import (
	"sync"
)

// req channelパターン
type Req struct {
	closing struct {
		ch chan struct{}
		sync.WaitGroup
		sync.Mutex
	}
	op1 struct {
		reqch chan<- string
	}
}

func NewReq() *Req {
	m := &Req{}
	m.closing.ch = make(chan struct{})

	op1reqch := make(chan string)
	m.op1.reqch = op1reqch

	m.closing.Add(1)
	go func() {
		m.op1Proc(op1reqch, m.closing.ch)
		m.closing.Done()
	}()

	return m
}

func (m *Req) Close() {
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

func (m *Req) Op1(val string) {
	select {
	case m.op1.reqch <- val:
		break
	case <-m.closing.ch:
		panic("its closed")
	}
}

func (*Req) op1Proc(in <-chan string, closingIn <-chan struct{}) {
	for {
		select {
		case op1 := <-in:
			// something todo
			_ = op1
		case <-closingIn:
			return
		}
	}
}
