package concurrentobjects

import (
	"fmt"
	"runtime"
	"sync"
)

type Model2 struct {
	closing struct {
		ch chan struct{}
		*sync.WaitGroup
	}
	prop1 struct {
		ch     <-chan string
		setch  chan<- string
		donech <-chan string
	}
}

func NewModel2() *Model2 {
	m := &Model2{}
	m.closing.ch = make(chan struct{})
	m.closing.WaitGroup = &sync.WaitGroup{}

	prop1ch, prop1setch, prop1donech := make(chan string), make(chan string), make(chan string)
	m.prop1.ch, m.prop1.setch, m.prop1.donech = prop1ch, prop1setch, prop1donech

	m.closing.Add(1)
	go func(closingWg *sync.WaitGroup, closingCh <-chan struct{}) {
		prop1Proc(prop1setch, closingCh, prop1ch, prop1donech)
		defer closingWg.Done()
	}(m.closing.WaitGroup, m.closing.ch)
	runtime.SetFinalizer(m, finalizeModel2)

	return m
}

func (m *Model2) Close() {
	close(m.closing.ch)
	m.closing.Wait()
}

func (m *Model2) Prop1() string {
	select {
	case val := <-m.prop1.ch:
		return val
	case <-m.closing.ch:
		panic("its closed")
	}
}

func (m *Model2) SetProp1(val string) {
	select {
	case m.prop1.setch <- val:
		nval := <-m.prop1.donech
		if val != nval {
			panic(fmt.Errorf("expect %s but %s", val, nval))
		}
	case <-m.closing.ch:
		panic("its closed")
	}
}

func prop1Proc(in <-chan string, closingIn <-chan struct{}, out chan<- string, doneOut chan<- string) {
	defer close(doneOut)
	defer close(out)

	var prop1 string
	for {
		select {
		case newProp := <-in:
			prop1 = newProp
			doneOut <- prop1
		case out <- prop1:
			break
		case <-closingIn:
			return
		}
	}
}

func finalizeModel2(m *Model2) {
	select {
	case <-m.closing.ch:
		break
	default:
		m.Close()
	}
}
