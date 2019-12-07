package main

import (
	"fmt"
	"time"
)

func main() {
}
func syncplay() {
	ch := make(chan string)

	go func() {
		time.Sleep(3)
		ch <- "ya"
	}()

	fmt.Println("hoge")

	str, ok := <-ch

	fmt.Println(ok)

	close(ch)

	fmt.Println(str)
}
