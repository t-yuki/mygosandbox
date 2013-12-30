package concurrentobjects_test

import (
	cobjs "."
	"fmt"
	"runtime"
	"testing"
)

type Model interface {
	Prop1() string
	SetProp1(string)
}

func TestModels(t *testing.T) {
	runtime.GC()
	orig := runtime.NumGoroutine()
	models := make([]Model, 0, 100)
	for i := 0; i < 10; i++ {
		models = append(models, cobjs.NewModel1())
		models = append(models, cobjs.NewModel2())
		models = append(models, cobjs.NewModel3())
	}

	for i := 0; i < 10000; i++ {
		go func(obj1 Model) {
			obj1.SetProp1(fmt.Sprintf("%d", i))
			obj1.Prop1()
		}(models[i%len(models)])
	}
	for i := 0; i < 10000; i++ {
		runtime.Gosched()
	}
	for i := range models {
		models[i] = nil
		runtime.GC()
	}
	runtime.GC()

	if runtime.NumGoroutine()-orig > 5 {
		t.Fatalf("expected goroutine: %d but %d", orig, runtime.NumGoroutine())
	}
}
