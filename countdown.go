package main

import (
	"os"
	"time"

	"github.com/basokant/go-with-tests/mocking"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, SleepFn: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
