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
	"strconv"
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
		
		// Check if white fill is requested via URL parameter
		useColor := r.URL.Query().Get("color") != "false"
		
		// Parse width and height from URL parameters (default: 600x320)
		canvasWidth := 600
		canvasHeight := 320
		if widthStr := r.URL.Query().Get("width"); widthStr != "" {
			if w, err := strconv.Atoi(widthStr); err == nil && w > 0 {
				canvasWidth = w
			}
		}
		if heightStr := r.URL.Query().Get("height"); heightStr != "" {
			if h, err := strconv.Atoi(heightStr); err == nil && h > 0 {
				canvasHeight = h
			}
		}
		
		// Generate and write surface SVG
		surface.Surface(w, useColor, canvasWidth, canvasHeight)
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
