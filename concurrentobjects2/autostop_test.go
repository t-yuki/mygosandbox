package concurrentobjects2_test

import (
	cobjs "."
	"bytes"
	"fmt"
	"runtime"
	"runtime/pprof"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestAutostop(t *testing.T) {
	func() {
		var wg sync.WaitGroup

		models := make([]*cobjs.Autostop, 0, 100)
		for i := 0; i < 3; i++ {
			models = append(models, cobjs.NewAutostop())
		}

		for i := 0; i < len(models); i++ {
			wg.Add(1)
			go func(m *cobjs.Autostop) {
				defer wg.Done()
				val := fmt.Sprintf("%d", i)
				if m.SetProp1(val); m.Prop1() != val {
					t.Errorf("expected %s but %s", val, m.Prop1())
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
		if strings.Contains(line, "concurrentobjects2.") {
			return lines, false
		}
	}
	return nil, true
}
