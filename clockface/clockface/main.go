package main

import (
	"os"
	"time"

	"github.com/basokant/go-with-tests/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
