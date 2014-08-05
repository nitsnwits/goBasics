package main

import "fmt"

func main() {
	const untyped = 123
	const typed int = 12345

	fmt.Printf("untyped 123 \t %T [%v]\n", untyped, untyped)
	fmt.Printf("typed 12345 \t %T [%v]\n", typed, typed)

	const a = "hello"
	const b = " world"
	const vartyped int = 3 * 4
	fmt.Printf("vartyped 3 * 4 \t %T [%v]\n", vartyped, vartyped)

	const err = a + b
	fmt.Printf("err \t %T [%v]\n", err, err)
}
