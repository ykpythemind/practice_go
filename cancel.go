package main

import (
	"context"
	"fmt"
	"time"
)

type work struct {
	id       int
	progress int
	err      error
	tick     *time.Ticker
}

func newWork(dur time.Duration, id int) *work {
	tick := time.NewTicker(dur)

	return &work{
		id:       id,
		progress: 0,
		tick:     tick,
	}
}

func (w *work) Do(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			w.stop()
			fmt.Printf("stop by cancel context! id: %d\n", w.id)
			return
		case <-w.tick.C:
			w.progress++

			str := ""
			for i := 0; i < w.progress; i++ {
				str += "*"
			}
			fmt.Printf("id:%d %s\n", w.id, str)

			if w.progress >= 100 {
				fmt.Printf("work id:%d complete\n", w.id)
				w.stop()
				return
			}
		}
	}
}

func (w *work) stop() {
	w.tick.Stop()
}

func cancelFunc() []*work {
	works := []*work{
		newWork(30*time.Millisecond, 1), newWork(98*time.Millisecond, 2), newWork(10*time.Millisecond, 3),
	}

	ctx, cancel := context.WithCancel(context.Background())

	for _, w := range works {
		go w.Do(ctx)
	}

	// n秒間回しまくる
	sec := 5
	time.Sleep(time.Duration(sec) * time.Second)

	// n秒たったら、キャンセルをctx経由で伝播させる
	cancel()

	for {
		select {
		case <-ctx.Done(): // キャンセル終わるまでここでブロックする
			fmt.Println("parent done", ctx.Err())

			fmt.Println("--- result ---")
			for _, w := range works {
				fmt.Printf("id: %d, progress: %d\n", w.id, w.progress)
			}

			return works
		}
	}
}
