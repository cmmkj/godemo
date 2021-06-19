package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, error error) {
	// startA := time.Now()
	// http.Get(a)
	// aDuration := time.Since(startA)

	// startB := time.Now()
	// http.Get(b)
	// bDuration := time.Since(startB)

	// if aDuration > bDuration {
	// 	return b
	// }
	// return a

	// aDuration := measureResponseTime(a)
	// bDuration := measureResponseTime(b)

	// if aDuration < bDuration {
	// 	return a
	// }
	// return b

	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Millisecond):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
