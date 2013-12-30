package goenumtest_test

import (
	e "."
	"testing"
)

func TestEnum1(t *testing.T) {
	val1 := e.Enum1B
	if val1 == e.Enum1A {
		t.Error()
	}

	// error, Enum1B is const
	// e.Enum1B = e.Enum1A

	// but creating with type cast is allowed so Enum1B == enum1c
	enum1c := e.Enum1(1)
	val2 := e.Enum1B
	if val2 != enum1c {
		t.Error()
	}
}

func TestEnum2(t *testing.T) {
	val1 := e.Enum2B
	if val1 == e.Enum2A {
		t.Error()
	}

	// huh?
	e.Enum2B = e.Enum2A

	// omg
	val2 := e.Enum2B
	if val2 != e.Enum2A {
		t.Error()
	}
}

func TestEnum3(t *testing.T) {
	val1 := e.Enum3B()
	if val1 == e.Enum3A() {
		t.Error()
	}

	// would you declare it?
	// e.SetEnum3B(e.Enum3A())

	// error, unexported
	// enum3c := e.Enum3{666}

	// but creating with default value 0 is allowed so Enum3A == enum3c.
	enum3c := e.Enum3{}
	val2 := e.Enum3A()
	if val2 != enum3c {
		t.Error()
	}
}
