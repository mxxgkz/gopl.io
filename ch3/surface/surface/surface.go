// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Package surface generates SVG rendering of 3-D surface functions.
package surface

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// Surface generates an SVG rendering of a 3-D surface function and writes it to out.
// If useColor is true, polygons are colored based on height (blue for valleys, red for peaks).
// If useColor is false, polygons are filled with white.
// If canvasWidth or canvasHeight are 0, default values (600x320) are used.
func Surface(out io.Writer, useColor bool, canvasWidth, canvasHeight int) {
	// Use defaults if not specified
	if canvasWidth <= 0 {
		canvasWidth = width
	}
	if canvasHeight <= 0 {
		canvasHeight = height
	}
	
	// Recalculate scales based on canvas dimensions
	xyscale := float64(canvasWidth) / 2 / xyrange
	zscale := float64(canvasHeight) * 0.4
	
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>", canvasWidth, canvasHeight)
	
	var minZ, maxZ float64
	if useColor {
		// First pass: find min and max z values for color normalization
		minZ, maxZ = math.Inf(1), math.Inf(-1)
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				// Calculate z at center of cell
				x := xyrange * (float64(i)/cells - 0.5 + 0.5/cells)
				y := xyrange * (float64(j)/cells - 0.5 + 0.5/cells)
				z := f(x, y)
				if z < minZ {
					minZ = z
				}
				if z > maxZ {
					maxZ = z
				}
			}
		}
	}
	
	// Second pass: generate polygons
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, canvasWidth, canvasHeight, xyscale, zscale)
			bx, by := corner(i, j, canvasWidth, canvasHeight, xyscale, zscale)
			cx, cy := corner(i, j+1, canvasWidth, canvasHeight, xyscale, zscale)
			dx, dy := corner(i+1, j+1, canvasWidth, canvasHeight, xyscale, zscale)
			
			var fillColor string
			if useColor {
				// Calculate z at center of cell for coloring
				x := xyrange * (float64(i)/cells - 0.5 + 0.5/cells)
				y := xyrange * (float64(j)/cells - 0.5 + 0.5/cells)
				z := f(x, y)
				// Map z to color (blue for valleys, red for peaks)
				fillColor = heightToColor(z, minZ, maxZ)
			} else {
				fillColor = "white"
			}
			
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, fillColor)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, canvasWidth, canvasHeight int, xyscale, zscale float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(canvasWidth)/2 + (x-y)*cos30*xyscale
	sy := float64(canvasHeight)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// heightToColor maps a height value to a color between blue (valleys) and red (peaks).
// z is normalized to [0, 1] range between minZ and maxZ, then interpolated between blue and red.
func heightToColor(z, minZ, maxZ float64) string {
	// Normalize z to [0, 1] range
	var normalized float64
	if maxZ > minZ {
		normalized = (z - minZ) / (maxZ - minZ)
	} else {
		normalized = 0.5 // if all values are the same, use middle color
	}
	
	// Interpolate between blue (#0000ff) and red (#ff0000)
	// normalized = 0 -> blue, normalized = 1 -> red
	red := int(normalized * 255)
	blue := int((1 - normalized) * 255)
	
	return fmt.Sprintf("#%02x00%02x", red, blue)
}

