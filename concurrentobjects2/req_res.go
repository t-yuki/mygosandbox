package concurrentobjects2

import (
	"sync"
)

// req/res channelパターン
type ReqRes struct {
	closing struct {
		ch chan struct{}
		sync.WaitGroup
		sync.Mutex
	}
	op1 struct {
		reqch chan<- string
		resch <-chan string
	}
}

func NewReqRes() *ReqRes {
	m := &ReqRes{}
	m.closing.ch = make(chan struct{})

	op1reqch, op1resch := make(chan string), make(chan string)
	m.op1.reqch = op1reqch
	m.op1.resch = op1resch

	m.closing.Add(1)
	go func() {
		m.op1Proc(op1reqch, m.closing.ch, op1resch)
		m.closing.Done()
	}()

	return m
}

func (m *ReqRes) Close() {
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

func (m *ReqRes) Op1(val string) {
	select {
	case m.op1.reqch <- val:
		<-m.op1.resch
	case <-m.closing.ch:
		panic("its closed")
	}
}

func (*ReqRes) op1Proc(in <-chan string, closingIn <-chan struct{}, resOut chan<- string) {
	if cap(in) != 0 {
		panic("it must be cap 0")
	}
	for {
		select {
		case op1 := <-in:
			// something todo
			resOut <- op1
		case <-closingIn:
			return
		}
	}
}
