package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan int)

	go func(in chan int) {
		val1 := <-in
		fmt.Println("GO: get value from channel ", val1)
		fmt.Println("GO: after read from chan")
	}(ch1)

	ch1 <- 43

	fmt.Println("MAIN: after put value to channel")

	fmt.Scanln()
}
