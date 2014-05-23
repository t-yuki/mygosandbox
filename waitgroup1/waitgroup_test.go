package waitgroup

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func TestReuseWaitGroup(t *testing.T) {
	for i := 0; i < 10000; i++ {
		testReuseWaitGroup(t)
	}
}

func testReuseWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var cnt int32
	// go 1.2 has spurious wakeup from WaitGroup.Wait bug
	// see: https://code.google.com/p/go/issues/detail?id=7734
	//
	// usual: G0 and G1 wake up G2
	// G2: wg.Add |-> wg.Wait  -> wg.Add |-> wg.Wait ->
	// G0:        |-> wg.Done -/         |           /
	// G1:                               |-> Done  -/
	//
	// unusual: G0 wake up G2 twice and G1 does nothing.
	// G2: wg.Add |-> wg.Wait  -> wg.Add |-> wg.Wait ->
	// G0:        |-> wg.Done -/---------------------/
	// G1:                               |-> Done -> x
	//
	// Note that, even if issue 7734 is fixed, G0 may wake up G2 twice.
	// But the time is the same to G1 Done call so such a difference is unobservable.
	for i := 0; i < 2; i++ {
		wg.Add(1)
		atomic.AddInt32(&cnt, int32(1))
		go func(i int) {
			runtime.Gosched()
			atomic.AddInt32(&cnt, -1)
			wg.Done()
		}(i)
		wg.Wait()
		if x := atomic.LoadInt32(&cnt); x != 0 {
			t.Fatal(x)
		}
	}
}
