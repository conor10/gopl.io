// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"net/http"
)

const (
	defaultWidth, defaultHeight = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = defaultWidth / 2 / xyrange // pixels per x or y unit
	zscale        = defaultHeight * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	red           = "#ff0000"
	blue          = "#0000ff"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)


func main() {
	if len(os.Args) > 1 && os.Args[1] == "server" {
		http.HandleFunc("/", hander)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
	} else {
		draw(os.Stdout, defaultHeight, defaultWidth)
	}
}

func hander(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Unable to process params: %v\n", err)
		return
	}

	width, height := defaultWidth, defaultHeight

	widthString := r.Form.Get("width")
	heightString := r.Form.Get("height")

	if widthString != "" {
		width, _ = strconv.Atoi(widthString)
	}

	if heightString != "" {
		height, _ = strconv.Atoi(heightString)
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	draw(w, height, width)
}

func draw(writer io.Writer, height, width int) {
	fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", defaultWidth, defaultHeight)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, height, width)
			bx, by := corner(i, j, height, width)
			cx, cy := corner(i, j+1, height, width)
			dx, dy := corner(i+1, j+1, height, width)

			// Ex 3.1
			if isValid(ax, ay, bx, by, cx, cy, dx, dy) {
				// Ex 3.3 TODO: Complete
				var colour string
				if by > dy {
					colour = red
				} else {
					colour = blue
				}
				fmt.Fprintf(writer, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, colour)
			}
		}
	}
	fmt.Fprintln(writer, "</svg>")
}

func corner(i, j int, height, width int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width) /2 + (x-y)*cos30*xyscale
	sy := float64(height) /2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func isValid(vals ...float64) bool {
	for _, val := range vals {
		if math.IsNaN(val) || math.IsInf(val, 0) {
			return false
		}
	}
	return true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
