package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	// fastURL := "http://www.facebook.com"
	// slowURL := "http://www.quii.co.uk"

	// want := fastURL
	// got := Racer(slowURL, fastURL)

	// if got != want {
	// 	t.Errorf("got '%s' want '%s'", got, want)
	// }

	// slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	time.Sleep(200 * time.Millisecond)
	// 	w.WriteHeader(http.StatusOK)
	// }))

	// fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// }))
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	// slowURL := slowServer.URL
	// fastURL := fastServer.URL

	// want := fastURL
	// got := Racer(slowURL, fastURL)

	// if got != want {
	// 	t.Errorf("got '%s', want '%s'", got, want)
	// }
	_, err := Racer(slowServer.URL, slowServer.URL)

	if err != nil {
		t.Errorf("expected an error but didn't get one")
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
