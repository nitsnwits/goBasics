package main

import "fmt"

type user struct {
	name   string
	age    int
	gender string
}

func main() {
	user1 :=
		user{
			name:   "neeraj",
			age:    28,
			gender: "male",
		}

	fmt.Println(user1)

	user2 := struct {
		name string
		age  int
	}{
		name: "dippi",
		age:  24,
	}

	fmt.Println(user2)
	fmt.Println(user2.name)
}
