package asm1

func Asm1(x int64) (ret int64)

//go:noescape
func Thresh(buf []byte, th byte)

func ThreshNoasm(buf []byte, th byte) {
	for i := 0; i < len(buf); i++ {
		if buf[i] > th {
			buf[i] = 1
		} else {
			buf[i] = 0
		}
	}
}
