package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 11
	ch2 <- 2
	ch2 <- 22

LOOP:
	for {
		select {
		case val := <-ch1:
			fmt.Println("get val from ch1", val)
		case val2 := <-ch2:
			fmt.Println("get val from ch2", val2)
		default:
			fmt.Println("default case")
			break LOOP
		}
	}

}
