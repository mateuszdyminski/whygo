package main

import (
	"fmt"
	"time"
)

func main() {
	done := time.NewTimer(time.Second * 10) // tworzymy channel, który zwróci wartość po 10 sekundach

	// nieskończona pętla
	for {
		select {
		case <-time.After(time.Second * 1): // wysyłamy informację do channel co sekundę
			fmt.Println("Hello Wroclaw")
		case <-done.C: // czekamy na wiadomość, żeby zakończyć nieskończoną pętlę
			fmt.Println("Quitting")
			return
		}
	}
}
