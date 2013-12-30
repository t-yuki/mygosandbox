package main_test

import (
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestRead(t *testing.T) {
	for _, v := range []int{math.MaxInt32, 100, 1, 0} {
		t.Logf("start %d", v)
		path := fmt.Sprintf("testrd%d.pipe", v)
		os.Remove(path)
		if err := syscall.Mkfifo(path, 0644); err != nil {
			t.Fatal(err)
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			f, err := os.OpenFile(path, os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			for i := 0; i < v; i++ {
				if n, err := f.Write([]byte("abc")); err != nil {
					t.Logf("wr: %v", err)
					break
				} else if n == 0 {
					t.Logf("wr: 0")
					break
				}
			}
			f.Close()
			wg.Done()
		}()

		f, err := os.OpenFile(path, os.O_RDONLY, 0644)
		if err != nil {
			t.Fatal(err)
		}

		done := make(chan struct{})
		go func() {
			buf := make([]byte, 512)
			for {
				if n, err := f.Read(buf); err != nil {
					t.Logf("rd: %v", err)
					break
				} else if n == 0 {
					t.Logf("rd: 0")
					break
				}
			}
			done <- struct{}{}
		}()
		time.Sleep(time.Millisecond * time.Duration(50))
		f.Close()

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("")
		}
		wg.Wait()
		os.Remove(path)
	}
}

func TestWrite(t *testing.T) {
	for _, v := range []int{math.MaxInt32, 100, 1, 0} {
		t.Logf("start %d", v)
		path := fmt.Sprintf("testwr%d.pipe", v)
		os.Remove(path)
		if err := syscall.Mkfifo(path, 0644); err != nil {
			t.Fatal(err)
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			f, err := os.OpenFile(path, os.O_RDONLY, 0644)
			if err != nil {
				panic(err)
			}

			buf := make([]byte, 100)
			for i := 0; i < v; i++ {
				if n, err := f.Read(buf); err != nil {
					t.Logf("rd: %v", err)
					break
				} else if n == 0 {
					t.Logf("rd: 0")
					break
				}
			}
			f.Close()
			wg.Done()
		}()

		done := make(chan struct{})
		f, err := os.OpenFile(path, os.O_WRONLY, 0644)
		if err != nil {
			t.Fatal(err)
		}
		go func() {
			for {
				if n, err := f.Write([]byte("abc")); err != nil {
					t.Logf("wr: %v", err)
					break
				} else if n == 0 {
					t.Logf("wr: 0")
					break
				}
			}
			done <- struct{}{}
		}()

		time.Sleep(time.Millisecond * time.Duration(50))
		f.Close()

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("")
		}
		wg.Wait()
		os.Remove(path)
	}
}

func TestReadFull(t *testing.T) {
	for _, v := range []int{math.MaxInt32, 1000, 1, 0} {
		t.Logf("start %d", v)
		path := fmt.Sprintf("testrd%d.pipe", v)
		os.Remove(path)
		if err := syscall.Mkfifo(path, 0644); err != nil {
			t.Fatal(err)
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			out, err := os.OpenFile(path, os.O_WRONLY, 0644)
			if err != nil {
				t.Fatal(err)
			}

			for i := 0; i < v; i++ {
				if n, err := out.Write([]byte("abc")); err != nil {
					t.Logf("wr: %v", err)
					break
				} else if n == 0 {
					t.Logf("wr: 0")
					break
				}
			}
			out.Close()
			wg.Done()
		}()

		in, err := os.OpenFile(path, os.O_RDONLY, 0644)
		if err != nil {
			t.Fatal(err)
		}

		done := make(chan struct{})
		go func() {
			buf := make([]byte, 512)
			for {
				if n, err := io.ReadFull(in, buf); err != nil {
					t.Logf("rd: %v", err)
					break
				} else if n == 0 {
					t.Logf("rd: 0")
					break
				}
			}
			done <- struct{}{}
		}()
		time.Sleep(time.Millisecond * time.Duration(50))

		in.Close()

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("")
		}
		wg.Wait()
		os.Remove(path)
	}
}

func TestDuplex(t *testing.T) {
	for _, v := range []int{math.MaxInt32, 1000, 1, 0} {
		t.Logf("start %d", v)
		path1 := fmt.Sprintf("testwr%d.pipe", v)
		path2 := fmt.Sprintf("testrd%d.pipe", v)
		os.Remove(path1)
		os.Remove(path2)
		if err := syscall.Mkfifo(path1, 0644); err != nil {
			t.Fatal(err)
		}
		if err := syscall.Mkfifo(path2, 0644); err != nil {
			t.Fatal(err)
		}

		var wg sync.WaitGroup
		wg.Add(1)

		cmd := exec.Command("./iomain.sh", "-path1", path2, "-path2", path1, "-count", fmt.Sprintf("%d", v))
		if err := cmd.Start(); err != nil {
			t.Fatal(err)
		}
		out, err := os.OpenFile(path1, os.O_WRONLY, 0644)

		go func() {
			if err != nil {
				panic(err)
			}
			for {
				if n, err := out.Write([]byte("abc")); err != nil {
					t.Logf("wr: %v", err)
					break
				} else if n == 0 {
					t.Logf("wr: 0")
					break
				}
			}
			wg.Done()
		}()
		done := make(chan struct{})
		go func() {
			in, err := os.OpenFile(path2, os.O_RDONLY, 0644)
			if err != nil {
				panic(err)
			}

			buf := make([]byte, 512)
			for {
				if n, err := io.ReadFull(in, buf); err != nil {
					t.Logf("rd: %v", err)
					break
				} else if n == 0 {
					t.Logf("rd: 0")
					break
				}
			}
			in.Close()
			done <- struct{}{}
		}()
		time.Sleep(time.Millisecond * time.Duration(50))
		out.Close()
		wg.Wait()
		// cmd.Wait()

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("")
		}
		os.Remove(path1)
		os.Remove(path2)
	}
}
