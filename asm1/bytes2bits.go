package asm1

//go:noescape
func BytesToBits(dst []byte, src []byte)

func BytesToBitsNoasm(dst []byte, src []byte) {
	for i := 0; i < len(dst); i++ {
		var b byte
		s := src[i*8]
		b |= (s >> 7) << 0

		s = src[i*8+1]
		b |= (s >> 7) << 1

		s = src[i*8+2]
		b |= (s >> 7) << 2

		s = src[i*8+3]
		b |= (s >> 7) << 3

		s = src[i*8+4]
		b |= (s >> 7) << 4

		s = src[i*8+5]
		b |= (s >> 7) << 5

		s = src[i*8+6]
		b |= (s >> 7) << 6

		s = src[i*8+7]
		b |= (s >> 7) << 7

		dst[i] = b
	}
}
