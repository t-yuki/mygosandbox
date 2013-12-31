package cobjs1

import (
	"runtime"
	"sync"
)

type ModelB1 struct {
	quit struct {
		ch chan struct{}
		*sync.WaitGroup
	}
	pingCh chan<- struct {
		val   string
		resch chan<- string
	}
}

func NewModelB1() *ModelB1 {
	m := &ModelB1{}
	m.quit.ch = make(chan struct{})
	m.quit.WaitGroup = &sync.WaitGroup{}

	pingCh := make(chan struct {
		val   string
		resch chan<- string
	}, 5)
	m.pingCh = pingCh

	m.quit.Add(1)
	go func(quitCh <-chan struct{}, quitWg *sync.WaitGroup) {
		modelB1PingProc(pingCh, quitCh)
		defer quitWg.Done()
	}(m.quit.ch, m.quit.WaitGroup)

	runtime.SetFinalizer(m, (*ModelB1).finalize)

	return m
}

func (m *ModelB1) Ping(val string) (pong string) {
	resch := make(chan string)
	p := struct {
		val   string
		resch chan<- string
	}{val, resch}
	m.pingCh <- p
	nval := <-resch
	return nval
}

func (m *ModelB1) finalize() {
	close(m.quit.ch)
	m.quit.Wait()
}

func modelB1PingProc(in <-chan struct {
	val   string
	resch chan<- string
}, quitIn <-chan struct{}) {
	for {
		select {
		case req := <-in:
			val := req.val
			// do something
			res := val
			req.resch <- res
		case <-quitIn:
			return
		}
	}
}
