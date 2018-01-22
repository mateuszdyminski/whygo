package main

import (
	"fmt"
	"time"
)

func main() {
	done := time.NewTimer(time.Second * 10) // create channel which puts value into it after 10s

	// infinite loop
	for {
		select {
		case <-time.After(time.Second * 1): // get value from channel each second
			fmt.Println("Hello Wroclaw")
		case <-done.C: // wait 10s until channel returns its value
			fmt.Println("Quitting")
			return
		}
	}
}
