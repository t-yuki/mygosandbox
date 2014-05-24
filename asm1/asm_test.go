package asm1

import "testing"

func TestAsm1(t *testing.T) {
	if n := Asm1(1); n != 2 {
		t.Fatal(n)
	}
}

func TestThresh(t *testing.T) {
	testThresh(t, Thresh)
}

func TestThreshNoasm(t *testing.T) {
	testThresh(t, ThreshNoasm)
}

func testThresh(t *testing.T, thresh func(b []byte, th uint8)) {
	buf := make([]byte, 4096)
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i % 256)
	}
	thresh(buf, 100)
	for i := 0; i < len(buf); i++ {
		if i%256 > 100 && buf[i] != 1 {
			t.Fatalf("%d: %d", i, buf[i])
		}
		if i%256 <= 100 && buf[i] != 0 {
			t.Fatalf("%d: %d", i, buf[i])
		}
	}
}

// go test -bench .
// BenchmarkThresh 2000000000               0.02 ns/op
// BenchmarkThreshNoasm    2000000000               0.43 ns/op
func BenchmarkThresh(b *testing.B) {
	buf := make([]byte, 4096*1024*64)
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i % 256)
	}
	b.ResetTimer()
	Thresh(buf, 100)
}

func BenchmarkThreshNoasm(b *testing.B) {
	buf := make([]byte, 4096*1024*64)
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i % 256)
	}
	b.ResetTimer()
	ThreshNoasm(buf, 100)
}
