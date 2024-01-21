// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	// "os"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Listening on: :80")

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

// handler sends a lissajous gif to http client
// and takes input from the url to control the generation of the lissajous figure
func handler(w http.ResponseWriter, r *http.Request) {
	// default gen parameters
	var (
		cycles  = 5.0     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 200   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 5     // delay between frames in 10ms units
	)
	var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xFF}}

	// get url
	fmt.Println(r.URL.String())


	lissajous(w, cycles, res, size, nframes, delay, palette)
}

// lissajous generates lissajous gif and writes it to out
func lissajous(out io.Writer, cycles, res float64, size, nframes, delay int, palette color.Palette) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
