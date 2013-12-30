package concurrentobjects

import (
	"fmt"
	"runtime"
	"sync"
)

type Model3 struct {
	closing struct {
		ch chan struct{}
		*sync.WaitGroup
	}
	prop1 struct {
		ch    <-chan string
		setch chan struct {
			val    string
			donech chan<- string
		}
	}
}

func NewModel3() *Model3 {
	m := &Model3{}
	m.closing.ch = make(chan struct{})
	m.closing.WaitGroup = &sync.WaitGroup{}

	prop1ch, prop1setch := make(chan string), make(chan struct {
		val    string
		donech chan<- string
	}, 5)
	m.prop1.ch = prop1ch
	m.prop1.setch = prop1setch

	m.closing.Add(1)
	go func(closing *sync.WaitGroup, closingCh <-chan struct{}) {
		model3prop1Proc(prop1setch, closingCh, prop1ch)
		defer closing.Done()
	}(m.closing.WaitGroup, m.closing.ch)
	runtime.SetFinalizer(m, finalizeModel3)

	return m
}

func (m *Model3) Close() {
	close(m.closing.ch)
	m.closing.Wait()
}

func (m *Model3) Prop1() string {
	return <-m.prop1.ch
}

func (m *Model3) SetProp1(val string) {
	var p struct {
		val    string
		donech chan<- string
	}
	donech := make(chan string)
	p.val = val
	p.donech = donech
	select {
	case m.prop1.setch <- p:
		nval := <-donech
		if val != nval {
			panic(fmt.Errorf("expect %s but %s", val, nval))
		}
	case <-m.closing.ch:
		panic("its closed")
	}
}

func model3prop1Proc(in <-chan struct {
	val    string
	donech chan<- string
}, closingIn <-chan struct{}, out chan<- string) {
	var prop1 string
	for {
		select {
		case p := <-in:
			prop1 = p.val
			if p.donech != nil {
				p.donech <- prop1
			}
		case out <- prop1:
			break
		case <-closingIn:
			return
		}
	}
}

func finalizeModel3(m *Model3) {
	select {
	case <-m.closing.ch:
	default:
		m.Close()
	}
}
