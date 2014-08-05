package main

import "fmt"

type animal struct {
	name     string
	category string
}

type dog struct {
	breed string
	barks bool
	animal
}

func main() {
	jack := dog{
		breed: "something",
		barks: true,
		animal: animal{
			name:     "Jack",
			category: "black",
		},
	}
	fmt.Println(jack)
}
