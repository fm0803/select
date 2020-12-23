package main

import (
	"fmt"
	 "time"
)

func main() {
	ch := make (chan int,10)
	defer close(ch)
	go func() {
		// do something
		time.Sleep(time.Second*3) // sleep one second
		ch <- 1
	}()
	select {
	case <- ch:
	case <- time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}
