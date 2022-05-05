//go:build ignore

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func main() {

	in, err := os.Open("temp/node_pf/node_modules/@patternfly/patternfly/assets/images/pf-c-brand__logo-on-sm.svg")
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	icon, _ := oksvg.ReadIconStream(in)

	sizes := []int{192, 512}
	for _, size := range sizes {
		w, h := size, size

		icon.SetTarget(0, 0, float64(w), float64(h))
		rgba := image.NewRGBA(image.Rect(0, 0, w, h))

		for x := 0; x < size; x++ {
			for y := 0; y < size; y++ {
				rgba.Set(x, y, color.NRGBA{R: 255, G: 255, B: 255, A: 255})
			}
		}

		icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)

		var out *os.File
		out, err = os.Create(fmt.Sprintf("server/web/logo-%d.png", size))
		if err != nil {
			log.Fatal(err)
		}

		err = png.Encode(out, rgba)
		if err != nil {
			log.Fatal(err)
		}
		err = out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

}
