package main

import (
	"testing"
)

func TestCancel(t *testing.T) {
	works := cancelFunc()

	w := works[0] // works[1] => fail
	if w.progress != 100 {
		t.Errorf("work id:%d must finished but progress is %d", w.id, w.progress)
	}
}
