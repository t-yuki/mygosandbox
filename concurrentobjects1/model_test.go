package cobjs1_test

import (
	cobjs "."
	"bytes"
	"fmt"
	"runtime"
	"runtime/pprof"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestModels_serial(t *testing.T) {
	func() {
		var wg sync.WaitGroup

		models := make([]interface{}, 0, 100)
		for i := 0; i < 3; i++ {
			models = append(models, cobjs.NewModel1())
			models = append(models, cobjs.NewModel2())
			models = append(models, cobjs.NewModel3())
			models = append(models, cobjs.NewModel4())
			models = append(models, cobjs.NewModelB1())
		}

		for i := 0; i < len(models); i++ {
			wg.Add(1)
			go func(obj1 interface{}) {
				defer wg.Done()
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
				if m, ok := obj1.(cobjs.ModelAExt2); ok {
					if m.SetProp2(i); m.Prop2() != i {
						t.Errorf("expected %d but %d", i, m.Prop2())
					}
				}
				if m, ok := obj1.(cobjs.ModelAExt1); ok {
					m.Close()
				}
			}(models[i])
		}

		// Wait to cleanup all goroutines.
		wg.Wait()
	}()

	forceGC()

	if lines, ok := checkGoroutineLeaks(); !ok {
		for _, l := range lines {
			t.Log(l)
		}
		t.Fatal("leaked goroutine detected")
	}
}

func TestModels_parallel(t *testing.T) {
	func() {
		var recoverCount int32
		var wg sync.WaitGroup

		models := make([]interface{}, 0, 100)
		for i := 0; i < 10; i++ {
			models = append(models, cobjs.NewModel1())
			models = append(models, cobjs.NewModel2())
			models = append(models, cobjs.NewModel3())
			models = append(models, cobjs.NewModel4())
			models = append(models, cobjs.NewModelB1())
		}

		for i := 0; i < 10000; i++ {
			wg.Add(1)
			go func(obj1 interface{}) {
				defer wg.Done()
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
				if m, ok := obj1.(cobjs.ModelAExt2); ok {
					m.SetProp2(i)
					m.Prop2()
				}
				if m, ok := obj1.(cobjs.ModelB); ok {
					if pong := m.Ping(val); pong != val {
						t.Errorf("expected %s but %s", val, pong)
					}
				}
			}(models[i%len(models)])
		}

		// Try to interrupt by Close
		for _, obj1 := range models {
			if m, ok := obj1.(cobjs.ModelAExt1); ok {
				m.Close()
			}
		}

		// Wait to cleanup all goroutines.
		wg.Wait()
		t.Logf("recovered: %d", recoverCount)
	}()

	forceGC()

	if lines, ok := checkGoroutineLeaks(); !ok {
		for _, l := range lines {
			t.Log(l)
		}
		t.Fatal("leaked goroutine detected")
	}
}

func forceGC() {
	// `10` is a empirical value...
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
		runtime.GC()
	}
}

func checkGoroutineLeaks() ([]string, bool) {
	buf := &bytes.Buffer{}
	pprof.Lookup("goroutine").WriteTo(buf, 1)
	lines := strings.Split(buf.String(), "\n")
	for _, line := range lines {
		if strings.Contains(line, "concurrentobjects1.") {
			return lines, false
		}
	}
	return nil, true
}
