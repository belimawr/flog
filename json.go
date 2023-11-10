package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func GenerateJson(option *Option) error {
	if option.Output == "" {
		option.Output = ":3000"
	}
	log.Printf("listening on %s", option.Output)
	return http.ListenAndServe(option.Output, newHandler(option))
}

func newHandler(option *Option) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.String())
		created := time.Now()

		logLine := NewLog(option.Format, created)
		if option.Format == "json" {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		}
		fmt.Fprintf(w, logLine)
	}

	return fn
}
