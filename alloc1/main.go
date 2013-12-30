// +build main

package main

import "fmt"

func main() {
	var a1 *int = call1()
	fmt.Printf("value = %d", *a1)
}

func call1() *int {
	var b1 int = 10
	return &b1
}
