package mocking

import (
	"fmt"
	"io"
	"iter"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := range countDown(countdownStart) {
		_, _ = fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	_, _ = fmt.Fprint(out, finalWord)
}

func countDown(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}
