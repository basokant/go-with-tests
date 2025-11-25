package concurrency

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

const defaultTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, defaultTimeout)
}

var ErrTimeout = errors.New("timed out waiting for race")

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", errors.Join(ErrTimeout, fmt.Errorf("race between %s and %s", a, b))
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
