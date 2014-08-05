package main

import "fmt"

type hobby struct {
	name string
}

func (h hobby) display() {
	fmt.Printf("Your hobby is: %s\n", h.name)
}

type player struct {
	name   string
	plays  string
	AtBats string
	Hits   int
}

func (p player) battingAverage() { //do it later

}

func main() {
	h1 := hobby{
		name: "playing guitar",
	}
	h1.display()

}
