package main

import (
	"./shapes"
	"image"
	"image/color"
	"log"
)

func main() {
	img := shapes.FilledImage(500, 300, image.White)

	fill := color.RGBA{200, 200, 200, 255}

	for i := 0; i < 10; i++ {
		width, height := 50+(20*i), 30+(10*i)

		rect, err := shapes.New("rectangle", shapes.Option{Fill: fill, Rect: image.Rect(0, 0, width, height), Filled: true})
		if err != nil {
			log.Fatal(err)
		}

		x := 10 + (20 * i)
		for j := i / 2; j >= 0; j-- {
			rect.Draw(img, x+j, (x/2)+j)
		}

		fill.R -= uint8(i * 5)
		fill.G = fill.R
		fill.B = fill.R
	}

	shapes.SaveImage(img, "rectangle.png")
}
