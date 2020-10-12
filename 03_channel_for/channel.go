package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan int)

	go func(out chan<- int) {
		for i := 0; i <= 10; i++ {
			fmt.Println("before:", i)
			out <- i
			fmt.Println("after:", i)
		}
		close(out)
		fmt.Println("generator finish")
	}(ch1)

	for i := range ch1 {
		fmt.Printf("%d get \n", i)
	}

	fmt.Scanln()
}
