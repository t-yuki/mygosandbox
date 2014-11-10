package cgo

import "testing"

// test passes on `go version go1.3.3 linux/amd64`
func TestCallGOJump(t *testing.T) {
	if ret := CallGOJump(); ret != 1 {
		t.Fatal("want: 1 but: ", ret)
	}
	if ret := CallGOJumpNest1(); ret != 1 {
		t.Fatal("want: 1 but: ", ret)
	}
	if ret := CallGOJumpStruct(); ret != 1 {
		t.Fatal("want: 1 but: ", ret)
	}
}
