package asm1

import "testing"

func TestAsm1(t *testing.T) {
	if n := Asm1(1); n != 2 {
		t.Fatal(n)
	}
}
