package main

import (
	"errors"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/coreyvan/healthcheck-example/health"
)

func main() {

	health.RegisterPeriodicFunc("minute_even", time.Second*15, currentMinuteEvenCheck)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func currentMinuteEvenCheck() error {
	m := time.Now().Minute()
	if m%2 == 0 {
		return errors.New("current minute is even")
	}
	return nil
}
