package concurrentobjects_test

import (
	cobjs "."
	"testing"
)

func TestModel1(t *testing.T) {
	obj1 := cobjs.NewModel1()

	if obj1.SetProp1("a"); obj1.Prop1() != "a" {
		t.Fatal(obj1.Prop1())
	}
}
