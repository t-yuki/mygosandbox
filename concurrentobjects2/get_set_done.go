package concurrentobjects2

import (
	"sync"
)

// get + set/done channelパターン
type GetSetDone struct {
	closing struct {
		ch chan struct{}
		sync.WaitGroup
		sync.Mutex
	}
	prop1 struct {
		ch     <-chan string
		setch  chan<- string
		donech <-chan struct{}
	}
}

func NewGetSetDone() *GetSetDone {
	m := &GetSetDone{}
	m.closing.ch = make(chan struct{})

	prop1ch, prop1setch, prop1donech := make(chan string), make(chan string), make(chan struct{})
	m.prop1.ch = prop1ch
	m.prop1.setch = prop1setch
	m.prop1.donech = prop1donech

	m.closing.Add(1)
	go func() {
		m.getsetProc(prop1setch, m.closing.ch, prop1ch, prop1donech)
		m.closing.Done()
	}()

	return m
}

func (m *GetSetDone) Close() {
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

func (m *GetSetDone) Prop1() string {
	select {
	case p := <-m.prop1.ch:
		return p
	case <-m.closing.ch:
		panic("its closed")
	}
}

func (m *GetSetDone) SetProp1(val string) {
	select {
	case m.prop1.setch <- val:
		<-m.prop1.donech
	case <-m.closing.ch:
		panic("its closed")
	}
}

func (*GetSetDone) getsetProc(in <-chan string, closingIn <-chan struct{}, out chan<- string, doneOut chan<- struct{}) {
	if cap(in) != 0 {
		panic("it must be cap 0")
	}
	var prop1 string
	for {
		select {
		case newProp := <-in:
			prop1 = newProp
			doneOut <- struct{}{}
		case out <- prop1:
			break
		case <-closingIn:
			return
		}
	}
}
