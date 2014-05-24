package asm1

import "testing"

func TestC1(t *testing.T) {
	if n := C1(1); n != 2 {
		t.Fatal(n)
	}
}
