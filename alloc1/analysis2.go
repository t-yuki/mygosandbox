// +build analysis2

package main

import "fmt"

type IFType interface {
	Func1() int
}

type StructType struct {
	f1 int
}

func (r *StructType) Func1() int {
	return r.f1
}

func Unmarshal1(a string) *StructType {
	return &StructType{}
}

func Unmarshal2(a string) IFType {
	return &StructType{}
}

func main() {
	var a *StructType = Unmarshal1("a")
	fmt.Println(a.Func1())
	var b IFType = Unmarshal2("b")
	fmt.Println(b.Func1())
}
