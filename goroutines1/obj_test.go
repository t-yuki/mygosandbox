package goroutines1

import (
	"testing"
)

type StartStopWaiter interface {
	Start()
	Stop()
	Wait()
}

type Obj1 struct {
	status Lifestage
	x      Goroutine
	y      Goroutine
	z      Goroutine
}

func (o *Obj1) Start() {
	// start processes
	if !o.status.SetStarting() {
		// illegal state
		return
	}
	o.x.Start0(o.xProc)
	o.y.Start(o.yProc)
	o.z.Start(func(stopIn <-chan struct{}) {
		o.zProc(0, stopIn)
	})

	o.status.SetStarted()
}

func (o *Obj1) Stop() {
	if !o.status.SetStopped() {
		// illegal state
		return
	}
	o.x.Stop()
	o.y.Stop()
	o.z.Stop()
}

func (o *Obj1) Wait() {
	if !o.status.Stopped() {
		// illegal state
		return
	}
	o.x.Wait()
	o.y.Wait()
	o.z.Wait()
}

func (o *Obj1) xProc() {
}

func (o *Obj1) yProc(<-chan struct{}) {
}

func (o *Obj1) zProc(x int, stop <-chan struct{}) {
	<-stop
	if !o.z.Stopped() {
		panic("oops")
	}
}

func (o *Obj1) SetParam1(int) {
	if !o.status.SetCreating() {
		// illegal state
		return
	}
	defer o.status.SetCreated()
}

func Test1(t *testing.T) {
	obj1 := &Obj1{}
	obj1.SetParam1(0)
	obj1.Start()
	obj1.Stop()
	obj1.Wait()
}
