package compressor

import (
	"image"

	"github.com/nfnt/resize"
)

func Resize(maxWidth uint, img image.Image) image.Image {
	imgWidth := uint(img.Bounds().Dx())
	return resize.Resize(min(maxWidth, imgWidth), 0, img, resize.Lanczos3)
}

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}
