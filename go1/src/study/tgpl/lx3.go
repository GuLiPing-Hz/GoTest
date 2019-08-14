package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"os"
)

var palette = []color.Color{color.White, color.Black,
	//练习1.5
	color.RGBA{0, 0xff, 0, 0xff},
	color.RGBA{0xff, 0, 0, 0xff},
}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
	greenIndex = 2
	redIndex
)

func main() {
	file, _ := os.Create("x.gif")
	defer file.Close()
	lissajous(file)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := 0.5 //rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t, j := 0.0, 0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			//img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
			//练习题1.5
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
			//练习题1.6
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(j%3+1))
			j++
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
