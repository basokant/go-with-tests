package main

import (
	"os"

	"github.com/basokant/go-with-tests/mocking"
)

func main() {
	mocking.Countdown(os.Stdout)
}
