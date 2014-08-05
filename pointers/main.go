package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	var count int
	count = 20
	fmt.Println("Address: ", &count, " Value: ", count)

	fmt.Println("\n Ex 2: \n")

	var c2 int
	c2 = 42
	var c2p *int //this can be also done as c2p := &c2
	// zero value for a pointer is nil reserved keyword in Go
	c2p = &c2
	fmt.Println("Address of pointer: ", &c2p, " Value of pointer ", c2p, " Value that the pointer points to: ", *c2p)

	fmt.Println("\n Ex 3: \n")
	user1 := user{
		name: "neeraj",
		age:  29,
	}
	fmt.Println(user1)
	inc(&user1.age)
	changeString(&user1.name)
	fmt.Println(user1)

}

func inc(count *int) {
	*count++
}

func changeString(name *string) {
	*name = *name + "-changed by Pointer"
}
