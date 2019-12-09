package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func trap() {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt)

	timer := time.NewTimer(5 * time.Second)

	count := 0
	var mux sync.Mutex

	for {
		select {
		case <-ch:
			mux.Lock()
			count++
			mux.Unlock()

			if count >= 3 {
				log.Println("more than 3 times called. exit...")
				goto endlabel
			}

		case <-timer.C:
			log.Println("timer")
			goto endlabel
		}
	}

endlabel:
	os.Exit(0)
}
