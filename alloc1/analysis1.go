// +build analysis1

package main

import "fmt"
import "github.com/t-yuki/goalloctest/pkg1"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	a := min(5, 10)
	fmt.Println(a)

	b := pkg1.Max(5, 10)
	fmt.Println(b)

	var c int
	var c1 pkg1.MaxerImpl
	c = c1.Max1(5, 10)
	fmt.Println(c)
	var c2 pkg1.MaxerImpl
	c = c2.Max2(5, 10)
	fmt.Println(c)
	var c3 pkg1.MaxerImpl
	c = c3.Max3(5, 10)
	fmt.Println(c)
	var c4 pkg1.MaxerImpl
	c = c4.Max4(5, 10)
	fmt.Println(c)
	var c5 pkg1.MaxerImpl
	c = c5.Max5(5, 10)
	fmt.Println(c)

	var d int
	d1 := &pkg1.MaxerImpl{}
	d = d1.Max1(5, 10)
	fmt.Println(d)
	d2 := &pkg1.MaxerImpl{}
	d = d2.Max2(5, 10)
	fmt.Println(d)
	d3 := &pkg1.MaxerImpl{}
	d = d3.Max3(5, 10)
	fmt.Println(d)
	d4 := &pkg1.MaxerImpl{}
	d = d4.Max4(5, 10)
	fmt.Println(d)
	d5 := &pkg1.MaxerImpl{}
	d = d5.Max5(5, 10)
	fmt.Println(d)

	var e int
	e1 := pkg1.NewMaxerImpl()
	e = e1.Max1(5, 10)
	fmt.Println(e)
	e2 := pkg1.NewMaxerImpl()
	e = e2.Max2(5, 10)
	fmt.Println(e)
	e3 := pkg1.NewMaxerImpl()
	e = e3.Max3(5, 10)
	fmt.Println(e)
	e4 := pkg1.NewMaxerImpl()
	e = e4.Max4(5, 10)
	fmt.Println(e)
	e5 := pkg1.NewMaxerImpl()
	e = e5.Max5(5, 10)
	fmt.Println(e)

	var f int
	f1 := pkg1.NewMaxer()
	f = f1.Max1(5, 10)
	fmt.Println(f)
	f2 := pkg1.NewMaxer()
	f = f2.Max2(5, 10)
	fmt.Println(f)
	f3 := pkg1.NewMaxer()
	f = f3.Max3(5, 10)
	fmt.Println(f)
	f4 := pkg1.NewMaxer()
	f = f4.Max4(5, 10)
	fmt.Println(f)
	f5 := pkg1.NewMaxer()
	f = f5.Max5(5, 10)
	fmt.Println(f)

}
