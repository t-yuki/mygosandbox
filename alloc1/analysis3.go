// +build analysis3

package main

import "fmt"

func alloc1(a string) []byte {
	b := make([]byte, 10)
	return b
}

func alloc2(a string) []byte {
	b := make([]byte, 10)
	return b[0:5]
}

func alloc3(a string) []byte {
	b := &[10]byte{}
	return b[:]
}

func alloc4(a string) []byte {
	b := new([10]byte)
	return b[:]
}

func main() {
	var a []byte = alloc1("a")
	printbytes(a)
	var b []byte = alloc2("b")
	printbytes(b)
	var c []byte = alloc3("c")
	printbytes(c)
	var d []byte = alloc4("d")
	printbytes(d)
}

func printbytes(a []byte) {
	l := len(a)
	fmt.Println("len:%d", l)
	for v := range a {
		fmt.Print(v)
	}
	fmt.Println("")
}
