package cobjs1

import (
	"fmt"
	"runtime"
	"sync"
)

type Model4 struct {
	quit struct {
		ch chan struct{}
		*sync.WaitGroup
	}
	prop1 struct {
		ch     <-chan string
		setch  chan<- string
		donech <-chan string
	}
}

func NewModel4() *Model4 {
	m := &Model4{}
	m.quit.ch = make(chan struct{})
	m.quit.WaitGroup = &sync.WaitGroup{}

	prop1ch, prop1setch, prop1donech := make(chan string), make(chan string), make(chan string)
	m.prop1.ch, m.prop1.setch, m.prop1.donech = prop1ch, prop1setch, prop1donech

	m.quit.Add(1)
	go func(quitCh <-chan struct{}, quitWg *sync.WaitGroup) {
		model4prop1Proc(prop1setch, quitCh, prop1ch, prop1donech)
		defer quitWg.Done()
	}(m.quit.ch, m.quit.WaitGroup)

	runtime.SetFinalizer(m, (*Model4).finalize)

	return m
}

func (m *Model4) Prop1() string {
	return <-m.prop1.ch
}

func (m *Model4) SetProp1(val string) {
	m.prop1.setch <- val
	nval := <-m.prop1.donech
	if val != nval {
		panic(fmt.Errorf("expect %s but %s", val, nval))
	}
}

func (m *Model4) finalize() {
	close(m.quit.ch)
	m.quit.Wait()
}

func model4prop1Proc(in <-chan string, quitIn <-chan struct{}, out chan<- string, doneOut chan<- string) {
	var prop1 string
	for {
		select {
		case newProp := <-in:
			prop1 = newProp
			doneOut <- prop1
		case out <- prop1:
			break
		case <-quitIn:
			return
		}
	}
}
