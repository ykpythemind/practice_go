package main

import (
	"fmt"
	"time"
)

func ticktuck() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("done!")
			return
		case t := <-ticker.C:
			fmt.Printf("tick %s\n", t)
		}
	}
}
