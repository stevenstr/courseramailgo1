package main

import (
	"fmt"
)

func main() {

	cancelChan := make(chan struct{})
	dataChan := make(chan int)

	go func(cancelCh chan struct{}, dataCh chan int) {
		val := 0
		for {
			select {
			case <-cancelCh:
				return
			case dataCh <- val:
				val++
			}
		}
	}(cancelChan, dataChan)

	for currentVal := range dataChan {
		fmt.Println("read", currentVal)

		if currentVal > 4 {
			cancelChan <- struct{}{}
			break
		}
	}
}
