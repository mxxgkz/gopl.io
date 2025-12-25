// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 20.
//!+

// Server5 is a web server that computes surfaces and writes SVG data to the client.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"gopl.io/ch3/surface/surface"
)

var mu sync.Mutex
var count int

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		
		// Set Content-Type header for SVG
		w.Header().Set("Content-Type", "image/svg+xml")
		
		// Generate and write surface SVG
		surface.Surface(w)
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

//!-
