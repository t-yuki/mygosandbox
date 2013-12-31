package cobjs1_test

import (
	cobjs "."
	"fmt"
	"runtime"
	"sync/atomic"
	"testing"
)

func TestModels_serial(t *testing.T) {
	models := make([]interface{}, 0, 100)
	models = append(models, cobjs.NewModel1())
	models = append(models, cobjs.NewModel2())
	models = append(models, cobjs.NewModel3())
	models = append(models, cobjs.NewModelB1())

	for i := 0; i < len(models); i++ {
		go func(obj1 interface{}) {
			val := fmt.Sprintf("%d", i)
			if m, ok := obj1.(cobjs.ModelA); ok {
				if m.SetProp1(val); m.Prop1() != val {
					t.Errorf("expected %s but %s", val, m.Prop1())
				}
			}
			if m, ok := obj1.(cobjs.ModelB); ok {
				if pong := m.Ping(val); pong != val {
					t.Errorf("expected %s but %s", val, pong)
				}
			}
			if m, ok := obj1.(cobjs.ModelAExt1); ok {
				m.Close()
			}
		}(models[i])
	}
}

func TestModels_parallel(t *testing.T) {
	runtime.GC()
	orig := runtime.NumGoroutine()
	var recoverCount int32

	models := make([]interface{}, 0, 100)
	for i := 0; i < 10; i++ {
		models = append(models, cobjs.NewModel1())
		models = append(models, cobjs.NewModel2())
		models = append(models, cobjs.NewModel3())
		models = append(models, cobjs.NewModelB1())
	}

	for i := 0; i < 10000; i++ {
		go func(obj1 interface{}) {
			defer func() {
				if err := recover(); err != nil {
					if err != "its closed" {
						panic(err)
					}
					atomic.AddInt32(&recoverCount, 1)
				}
			}()
			val := fmt.Sprintf("%d", i)
			if m, ok := obj1.(cobjs.ModelA); ok {
				m.SetProp1(val)
				m.Prop1()
			}
			if m, ok := obj1.(cobjs.ModelB); ok {
				if pong := m.Ping(val); pong != val {
					t.Errorf("expected %s but %s", val, pong)
				}
			}
		}(models[i%len(models)])
	}

	for _, obj1 := range models {
		if m, ok := obj1.(cobjs.ModelAExt1); ok {
			runtime.Gosched()
			m.Close()
		}
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
	t.Logf("recovered: %d", recoverCount)
}
