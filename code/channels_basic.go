package main

import (
	"fmt"
)

func main() {
	done := make(chan bool) // tworzymy channel
	go func() {
		fmt.Println("Hello Wroclaw")
		done <- true // wysyłamy informacje, że nasza goroutine skończyła już prace
	}()

	<-done // czekamy, aż coś pojawi się w channel
}