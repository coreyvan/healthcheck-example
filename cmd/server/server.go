package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/coreyvan/healthcheck-example/health"
)

func main() {

	health.RegisterPeriodicFunc("minute_even", time.Second*15, func() error { return currentMinuteDivisible(2) })
	health.RegisterPeriodicFunc("minute_four", time.Second*15, func() error { return currentMinuteDivisible(4) })
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func currentMinuteDivisible(n int) error {
	m := time.Now().Minute()
	if m%n == 0 {
		return fmt.Errorf("current minute is divisible by %d", n)
	}
	return nil
}
