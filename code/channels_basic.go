package main

import (
	"fmt"
)

func main() {
	done := make(chan bool) // create channel
	go func() {
		fmt.Println("Hello Wroclaw")
		done <- true // send information that our goroutine finished its work
	}()

	<-done // wait until something occurs in channel
}