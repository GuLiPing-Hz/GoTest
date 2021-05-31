package main

import (
	"math/rand"
	"image/gif"
	"image"
	"math"
	"io"
	"net/http"
	"fmt"
	"log"
	"strconv"
	"image/color"
)

var palette2 = []color.Color{color.White, color.Black,
	color.RGBA{0, 0xff, 0, 0xff},
	color.RGBA{0xff, 0, 0, 0xff},
}

func lissajous2(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette2)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func gifer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cyclesStr := r.Form.Get("cycles")
	cycles, _ := strconv.Atoi(cyclesStr)

	if cycles <= 0 {
		cycles = 5
	}
	lissajous2(w, cycles)
}

func main() {
	//练习题1.12
	http.HandleFunc("/", handler)
	http.HandleFunc("/gif", gifer)
	http.ListenAndServe(":8000", nil)
}
