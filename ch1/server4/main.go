// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 20.
//!+

// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"gopl.io/ch1/lissajous/lissajous"
)

var mu sync.Mutex
var count int

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		
		// Parse cycles parameter from URL, default to 5
		cycles := 5
		if val := r.URL.Query().Get("cycles"); val != "" {
			if parsed, err := strconv.Atoi(val); err == nil {
				cycles = parsed
			}
		}
		
		lissajous.Lissajous(w, cycles)
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
