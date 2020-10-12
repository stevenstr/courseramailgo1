package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan int, 1)
	ch2 := make(chan int)

	//ch1 <- 1
	select {
	case val := <-ch1:
		fmt.Println("get val from ch1", val)
	case ch2 <- 2:
		fmt.Println("put to ch2")
	default:
		fmt.Println("default case")
	}

	fmt.Scanln()
}
