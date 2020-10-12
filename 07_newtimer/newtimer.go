package main

import (
	"fmt"
	"time"
)

func longSQLQuery() chan bool {
	ch := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- true
	}()
	return ch
}

func main() {

	timer := time.NewTimer(3 * time.Second)

	select {
	// if timer = 1 or 2
	case <-timer.C:
		fmt.Println("timer.C timeout is happend")
	case <-time.After(time.Minute):
		fmt.Println("time.After timeout is happend")
	//if timer = 3
	case result := <-longSQLQuery():
		if !timer.Stop() {
			<-timer.C
		}
		fmt.Println("operation result", result)
	}

}
