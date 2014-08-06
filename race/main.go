// Program uses a function type, closures and goroutines to
// calculate Fibonacci numbers. This program contains a Data Race.
package main

import (
	"fmt"
	"sync"
)

// fib is a function type that takes no parameters
// and returns an integer.
type fib func() int

// wg provides a mechanism to wait for goroutines
// to complete.
var wg sync.WaitGroup
var m sync.Mutex

// main is the entry point for all Go programs.
func main() {
	wg.Add(5)
	calculate(5, fibonacci())

	wg.Wait()
}

// fib returns a function that returns successive
// Fibonacci numbers. This function uses closure
// to maintain the state for the next call.
func fibonacci() fib {
	first, second := 0, 1
	return func() int {
		first, second = second, first+second
		return first
	}
}

// calculate recurses the call to the fibonacci function.
func calculate(toIndex int, f fib) {
	for i := 0; i < toIndex; i++ {
		go func() {
			m.Lock() {
				fmt.Println(f())
			}
			m.Unlock()
			wg.Done()
		}()
	}
}

