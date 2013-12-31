package propchan

import (
	"fmt"
)

type Property struct {
	ch     <-chan interface{}
	setch  chan<- interface{}
	donech <-chan interface{}
}

type Channels struct {
	In   <-chan interface{}
	Out  chan<- interface{}
	Done chan<- interface{}
}

func Make() (p Property, c Channels) {
	ch, setch, donech := make(chan interface{}), make(chan interface{}), make(chan interface{})
	p.ch, p.setch, p.donech = ch, setch, donech
	c.Out, c.In, c.Done = ch, setch, donech
	return
}

func (p *Property) Get() interface{} {
	return <-p.ch
}

func (p *Property) Set(val interface{}) {
	p.setch <- val
	nval := <-p.donech
	if val != nval {
		panic(fmt.Errorf("expect %s but %s", val, nval))
	}
}
