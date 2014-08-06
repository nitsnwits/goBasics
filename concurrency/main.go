package main

import (
	"fmt"
	"sync"
)

type fib func() int

var wg sync.WaitGroup

func main() {
	wg.Add(5)
	calculate(5, fibonacci())

	wg.Wait()
}

func fibonacci() fib {
	first, second := 0, 1
	return func() int {
		first, second = second, first+second
		return first
	}
}

func calculate(toIndex int, f fib) {
	for i := 0; i < toIndex; i++ {
		go func() {
			fmt.Println(f())
			wg.Done()
		}()
	}
}
