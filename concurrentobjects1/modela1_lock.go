package cobjs1

import (
	"sync"
)

type Model1 struct {
	prop1 struct {
		val string
		sync.Mutex
	}
}

func NewModel1() *Model1 {
	m := &Model1{}
	return m
}

func (m *Model1) Prop1() string {
	m.prop1.Lock()
	defer m.prop1.Unlock()
	return m.prop1.val
}

func (m *Model1) SetProp1(val string) {
	m.prop1.Lock()
	m.prop1.val = val
	m.prop1.Unlock()
}
