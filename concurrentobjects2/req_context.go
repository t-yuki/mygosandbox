package concurrentobjects2

import (
	"sync"
	"time"
)

// req context channelパターン
type ReqContext struct {
	closing struct {
		ch chan struct{}
		sync.WaitGroup
		sync.Mutex
	}
	op1 struct {
		reqch chan<- op1context
	}
}

type op1context struct {
	s        string
	resch    chan<- bool
	cancelch <-chan struct{}
}

// or done signal only
type op1context2 struct {
	s      string
	donech chan<- struct{}
}

func NewReqContext() *ReqContext {
	m := &ReqContext{}
	m.closing.ch = make(chan struct{})

	op1reqch := make(chan op1context)
	m.op1.reqch = op1reqch

	m.closing.Add(1)
	go func() {
		m.op1Proc(op1reqch, m.closing.ch)
		m.closing.Done()
	}()

	return m
}

func (m *ReqContext) Close() {
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

func (m *ReqContext) Op1(val string, cancel <-chan struct{}, res chan<- bool) {
	select {
	case m.op1.reqch <- op1context{s: val, cancelch: cancel, resch: res}:
		break
	case <-m.closing.ch:
		panic("its closed")
	}
}

func (*ReqContext) op1Proc(in <-chan op1context, closingIn <-chan struct{}) {
	for {
		select {
		case ctx := <-in:
			// something todo
			select {
			case <-ctx.cancelch:
				// check closingIn or not?
				ctx.resch <- false
				continue
			case <-time.After(time.Second):
				break
			}
			if ctx.resch != nil {
				ctx.resch <- true
			}
		case <-closingIn:
			return
		}
	}
}
