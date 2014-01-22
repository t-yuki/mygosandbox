package concurrentobjects2

import (
	"sync"
)

// req/done channelパターン
type ReqDone struct {
	closing struct {
		ch chan struct{}
		sync.WaitGroup
		sync.Mutex
	}
	op1 struct {
		reqch  chan<- string
		donech <-chan struct{}
	}
}

func NewReqDone() *ReqDone {
	m := &ReqDone{}
	m.closing.ch = make(chan struct{})

	op1reqch, op1donech := make(chan string), make(chan struct{})
	m.op1.reqch = op1reqch
	m.op1.donech = op1donech

	m.closing.Add(1)
	go func() {
		m.op1Proc(op1reqch, m.closing.ch, op1donech)
		m.closing.Done()
	}()

	return m
}

func (m *ReqDone) Close() {
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

func (m *ReqDone) Op1(val string) {
	select {
	case m.op1.reqch <- val:
		<-m.op1.donech
	case <-m.closing.ch:
		panic("its closed")
	}
}

func (*ReqDone) op1Proc(in <-chan string, closingIn <-chan struct{}, doneOut chan<- struct{}) {
	if cap(in) != 0 {
		panic("it must be cap 0")
	}
	for {
		select {
		case op1 := <-in:
			// something todo
			_ = op1
			doneOut <- struct{}{}
		case <-closingIn:
			return
		}
	}
}
