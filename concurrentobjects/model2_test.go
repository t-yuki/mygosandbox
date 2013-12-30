package concurrentobjects_test

import (
	cobjs "."
	"runtime"
	"sync"
	"testing"
)

func TestModel2(t *testing.T) {
	runtime.GC()
	orig := runtime.NumGoroutine()

	obj1 := cobjs.NewModel2()

	if obj1.SetProp1("a"); obj1.Prop1() != "a" {
		t.Fatal(obj1.Prop1())
	}

	obj1.Close()

	if runtime.NumGoroutine() != orig {
		t.Fatalf("expected goroutine: %d but %d", orig, runtime.NumGoroutine())
	}
}

func TestModel1Finalizer(t *testing.T) {
	runtime.GC()
	orig := runtime.NumGoroutine()

	for i := 0; i < 10; i++ {
		func1()
		if runtime.NumGoroutine() == orig {
			t.Fatalf("expected not equal: %d but %d", orig, runtime.NumGoroutine())
		}
		runtime.GC()
	}
	runtime.GC()
	runtime.Gosched()
	runtime.GC()

	if runtime.NumGoroutine() != orig {
		t.Fatalf("expected goroutine: %d but %d", orig, runtime.NumGoroutine())
	}
}

func func1() {
	var m sync.Mutex
	m.Lock()
	obj1 := cobjs.NewModel2()
	_ = obj1
}
