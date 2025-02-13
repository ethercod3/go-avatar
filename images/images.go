package images

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/ethercod3/goAvatar/colors"
	"github.com/ethercod3/goAvatar/patterns"
)

type boxOptions struct {
	baseXOffset int
	grid        []int
	boxSize     int
	img         *image.RGBA
	foreground  color.RGBA
}

func drawBoxes(b boxOptions) {
	yOffset := 0
	for _, cellIndex := range b.grid {
		x0 := (cellIndex * b.boxSize) + b.baseXOffset
		x1 := x0 + b.boxSize
		y0 := yOffset
		y1 := yOffset + b.boxSize
		rect := image.Rect(x0, y0, x1, y1)
		if cellIndex >= 0 {
			draw.Draw(b.img, rect, &image.Uniform{b.foreground}, image.Point{}, draw.Src)
		}
		yOffset += b.boxSize
	}
}

func Draw(p patterns.Pattern, scheme colors.ColorScheme, imageSizePx int, dimensions int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, imageSizePx, imageSizePx))
	background := colors.ColorToRGBA(scheme.First)
	draw.Draw(img, img.Bounds(), &image.Uniform{background}, image.Point{}, draw.Src)

	foreground := colors.ColorToRGBA(scheme.Second)
	boxSize := imageSizePx / (dimensions)

	options := boxOptions{
		baseXOffset: 0,
		grid:        p.Left,
		boxSize:     boxSize,
		img:         img,
		foreground:  foreground,
	}

	drawBoxes(options)
	options.baseXOffset += boxSize * (((dimensions - 1) / 2) - 1)

	if dimensions%2 != 0 {
		options.grid = p.Middle
		drawBoxes(options)
	}

	options.baseXOffset += boxSize * 2
	options.grid = p.Right
	drawBoxes(options)

	return img
}
