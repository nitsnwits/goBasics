package main

import "fmt"

type counter int

const c1 = 45

func main() {
	var v1 counter
	v1 = 0
	fmt.Printf("Value: %v Type: %T \n", v1, v1)

	fmt.Println("\n Ex 2: \n")

	var value int
	value = 10
	var v2 counter
	v2 = counter(value)
	fmt.Printf("Value: %v Type: %T \n", v2, v2)

	var v3 counter
	v3 = c1
	fmt.Printf("Value %v Type: %T\n", v3, v3)
}
