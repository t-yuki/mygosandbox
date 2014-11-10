package cgo

// #include "cgo.h"
import "C"

func CallGOJump() int {
	ret, _ := C.GOjump()
	return int(ret)
}

func CallGOJumpNest1() int {
	ret, _ := C.GOjump_nest1()
	return int(ret)
}

func CallGOJumpStruct() int {
	buf := &C.struct_GObuf{}
	ret, _ := C.GOjump_struct(buf)
	return int(ret)
}
