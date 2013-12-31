package cobjs1

import (
	"runtime"
	"sync"

	"github.com/t-yuki/mygosandbox/concurrentobjects1/propchan"
)

var _ ModelAExt2 = NewModel4()

type Model4 struct {
	quit struct {
		ch chan struct{}
		*sync.WaitGroup
	}
	prop1 propchan.Property
	prop2 propchan.Property
}

func NewModel4() *Model4 {
	m := &Model4{}
	m.quit.ch = make(chan struct{})
	m.quit.WaitGroup = &sync.WaitGroup{}

	var prop1chans, prop2chans propchan.Channels
	m.prop1, prop1chans = propchan.Make()
	m.prop2, prop2chans = propchan.Make()

	m.quit.Add(1)
	go func(quitCh <-chan struct{}, quitWg *sync.WaitGroup) {
		model4Proc(prop1chans, prop2chans, quitCh)
		defer quitWg.Done()
	}(m.quit.ch, m.quit.WaitGroup)

	runtime.SetFinalizer(m, (*Model4).finalize)

	return m
}

func (m *Model4) Prop1() string {
	return m.prop1.Get().(string)
}

func (m *Model4) SetProp1(val string) {
	m.prop1.Set(val)
}

func (m *Model4) Prop2() int {
	return m.prop2.Get().(int)
}

func (m *Model4) SetProp2(val int) {
	m.prop2.Set(val)
}

func (m *Model4) finalize() {
	close(m.quit.ch)
	m.quit.Wait()
}

func model4Proc(prop1chans propchan.Channels, prop2chans propchan.Channels, quitIn <-chan struct{}) {
	var prop1 string
	var prop2 int
	for {
		select {
		case newProp := <-prop1chans.In:
			prop1 = newProp.(string)
			prop1chans.Done <- prop1
		case newProp := <-prop2chans.In:
			prop2 = newProp.(int)
			prop2chans.Done <- prop2
		case prop1chans.Out <- prop1:
			continue
		case prop2chans.Out <- prop2:
			continue
		case <-quitIn:
			return
		}
	}
}
