package main

import (
	"fmt"
)

const (
	iterations = 7
	goroutines = 5
)

func doSomeWork(in int) {
	for j := 0; j < iterations; j++ {
		fmt.Println("goroutine:", goroutines-in, "iteration:", j)
		//runtime.Gosched()
	}
}

func main() {
	for i := 0; i < goroutines; i++ {
		go doSomeWork(i)
	}
	fmt.Scanln()
}
