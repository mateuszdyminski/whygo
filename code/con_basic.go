package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("Hello from goroutine") // HL
	time.Sleep(1 * time.Second)
}
