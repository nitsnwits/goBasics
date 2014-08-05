package main

import "fmt"

func main() {
	var intSlice []int
	counter := 10
	for i := 0; i <= 4; i++ {
		counter = counter + 10
		intSlice = append(intSlice, counter)
	}
	for _, val := range intSlice {
		fmt.Printf("Value: %d \n", val)
	}

	stringSlice := [5]string{"a", "b", "c", "d", "e"}
	slice1 := stringSlice[0:3]

	for _, val := range slice1 { //get more into iteration of slices for index as well as val
		fmt.Printf("Value: %s\n", val)
	}
}
