package concurrentobjects2

import (
	"runtime"
	"sync"
)

// Autostop パターン
type Autostop struct {
	quit struct {
		ch chan struct{}
		*sync.WaitGroup
	}
	prop1 struct {
		ch     <-chan string
		setch  chan<- string
		donech <-chan struct{}
	}
}

func NewAutostop() *Autostop {
	m := &Autostop{}
	m.quit.ch = make(chan struct{})
	m.quit.WaitGroup = &sync.WaitGroup{}

	prop1ch, prop1setch, prop1donech := make(chan string), make(chan string), make(chan struct{})
	m.prop1.ch, m.prop1.setch, m.prop1.donech = prop1ch, prop1setch, prop1donech

	m.quit.Add(1)
	go func(quitCh <-chan struct{}, quitWg *sync.WaitGroup) {
		autostopProp1Proc(prop1setch, quitCh, prop1ch, prop1donech)
		defer quitWg.Done()
	}(m.quit.ch, m.quit.WaitGroup)

	runtime.SetFinalizer(m, (*Autostop).finalize)

	return m
}

func (m *Autostop) Prop1() string {
	return <-m.prop1.ch
}

func (m *Autostop) SetProp1(val string) {
	m.prop1.setch <- val
	<-m.prop1.donech
}

func (m *Autostop) finalize() {
	close(m.quit.ch)
	m.quit.Wait()
}

func autostopProp1Proc(in <-chan string, quitIn <-chan struct{}, out chan<- string, doneOut chan<- struct{}) {
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
		case <-quitIn:
			return
		}
	}
}
