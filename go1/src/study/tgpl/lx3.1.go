package main

import (
	"math"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	//width, height = 600, 320            // canvas size in pixels
	cells   = 100  // number of grid cells
	xyrange = 30.0 // axis ranges (-xyrange..+xyrange)
	//xyscale       = width / 2 / xyrange // pixels per x or y unit
	//zscale        = height * 0.4        // pixels per z unit
	angle = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	//练习题3.4
	http.HandleFunc("/svg", svgHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("err=%s\n", err.Error())
	}
}

func svgHandler(w http.ResponseWriter, r *http.Request) {
	theMin, theMax := calMinMax()

	r.ParseForm()
	strWidth := r.Form.Get("width")
	strHeight := r.Form.Get("height")

	width, err := strconv.Atoi(strWidth)
	if err != nil || width == 0 {
		width = 600
	}
	height, err := strconv.Atoi(strHeight)
	if err != nil || height == 0 {
		height = 320
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	draw(w, theMin, theMax, width, height)
}

func draw(w io.Writer, min, max float64, width, height int) {
	l := max - min
	xyscale := float64(width) / 2 / xyrange // pixels per x or y unit
	zscale := float64(height) * 0.4         // pixels per z unit

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	wf, hf := float64(width), float64(height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok, _ := corner(i+1, j, wf, hf, xyscale, zscale)
			if !ok {
				continue
			}
			bx, by, ok, z := corner(i, j, wf, hf, xyscale, zscale)
			if !ok {
				continue
			}
			cx, cy, ok, _ := corner(i, j+1, wf, hf, xyscale, zscale)
			if !ok {
				continue
			}
			dx, dy, ok, _ := corner(i+1, j+1, wf, hf, xyscale, zscale)
			if !ok {
				continue
			}

			//练习题3.3
			r := uint8((z - min) / l * 255)
			col := fmt.Sprintf("#%02x00%02x", r, 255-r)
			//fmt.Println(col)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: %s;'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, col)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func calMinMax() (float64, float64) {
	var min, max float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5) // Compute surface height z.
			if z, ok := f(x, y); ok { // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
				if z < min {
					min = z
				}
				if z > max {
					max = z
				}
			}
		}
	}
	return min, max
}

func corner(i, j int, width, height, xyscale, zscale float64) (float64, float64, bool, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5) // Compute surface height z.
	z, ok := f(x, y)                        // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	if ! ok {
		return 0, 0, false, 0
	}
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true, z
}
func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	//练习题3.1
	if r == 0 {
		return 0, false
	}

	return math.Sin(r) / r, true
	//练习题3.2
	//return 0.2 * (math.Cos(x) + math.Cos(y)), true

	//a := 25.0
	//b := 17.0
	//a2 := a * a
	//b2 := b * b
	//r = y*y/a2 - x*x/b2
	//return r, true
}
