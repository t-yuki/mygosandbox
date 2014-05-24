package asm1

import "testing"

func TestBytesToBits(t *testing.T) {
	testBytesToBits(t, BytesToBits)
}

func TestBytesToBitsNoasm(t *testing.T) {
	testBytesToBits(t, BytesToBitsNoasm)
}

func testBytesToBits(t *testing.T, bitsToBytes func(dst []byte, src []byte)) {
	src := make([]byte, 4096*8)
	dst := make([]byte, 4096)

	for i := 0; i < len(src); i++ {
		src[i] = byte((i % 2) * 255)
	}
	bitsToBytes(dst, src)
	for i := 0; i < len(dst); i++ {
		if dst[i] != 170 {
			t.Fatalf("%d: %d", i, dst[i])
		}
	}

	for i := 0; i < len(src); i++ {
		switch i % 8 {
		case 0, 6:
			src[i] = 255
		default:
			src[i] = 0
		}
	}
	bitsToBytes(dst, src)
	for i := 0; i < len(dst); i++ {
		if dst[i] != 65 {
			t.Fatalf("%d: %d", i, dst[i])
		}
	}
}

// go test -bench .
// BenchmarkBytesToBits    2000000000               0.01 ns/op
// BenchmarkBytesToBitsNoasm       2000000000               0.17 ns/op
func BenchmarkBytesToBits(b *testing.B) {
	src := make([]byte, 4096*8*1024*4)
	dst := make([]byte, 4096*1024*4)

	for i := 0; i < len(src); i++ {
		src[i] = byte((i % 2) * 255)
	}
	b.ResetTimer()
	BytesToBits(dst, src)
}

func BenchmarkBytesToBitsNoasm(b *testing.B) {
	src := make([]byte, 4096*8*1024*4)
	dst := make([]byte, 4096*1024*4)

	for i := 0; i < len(src); i++ {
		src[i] = byte((i % 2) * 255)
	}
	b.ResetTimer()
	BytesToBitsNoasm(dst, src)
}
