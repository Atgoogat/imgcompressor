package compressor

import (
	"image"

	"github.com/nfnt/resize"
)

const resizeAlgo = resize.Lanczos3

func Resize(maxWidth uint, maxHeight uint, img image.Image) image.Image {
	if maxWidth != 0 && maxHeight != 0 {
		return resize.Thumbnail(maxWidth, maxHeight, img, resizeAlgo)
	}

	imgW, imgH := img.Bounds().Dx(), img.Bounds().Dy()
	return resize.Resize(min(uint(imgW), maxWidth), min(uint(imgH), maxHeight), img, resizeAlgo)
}

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}
