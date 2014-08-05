package main

import "fmt"

type (
	Speaker interface {
		SayHello()
	}

	English struct {
	}

	Chinese struct {
	}
)

func (s *English) SayHello() {
	fmt.Printf("Hello World!\n")
}

func (s *Chinese) SayHello() {
	fmt.Printf("你好世界\n")
}

func SayHello(s Speaker) {
	s.SayHello()
}

func main() {
	var speaker Speaker
	speaker = &English{} //speaker = new(English) -> creates a new pointer mem allocation and assigns the address to variable)
	speaker.SayHello()
	speaker = &Chinese{}
	speaker.SayHello()
	fmt.Println("\n Ex 2: \n")
	var s1 = new(English)
	var s2 = new(Chinese)
	SayHello(s1)
	SayHello(s2)
}
