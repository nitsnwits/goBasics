package main

import "fmt"

func main() {
	var a int
	var b string
	var c bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c bool \t %T [%v]\n", c, c)

	aa := "neeraj"
	bb := 10
	cc := true

	fmt.Printf("var a int \t %T [%v]\n", aa, aa)
	fmt.Printf("var b string \t %T [%v]\n", bb, bb)
	fmt.Printf("var c bool \t %T [%v]\n", cc, cc)

	aaa := float32(3.14)
	fmt.Printf("var aaa float32 \t %T [%v]\n", aaa, aaa)
}
