package compressor

import (
	"image"
	"testing"
)

func TestResize_KeepAspect(t *testing.T) {
	img := image.NewGray(image.Rect(0, 0, 100, 100))
	newImg := Resize(50, 40, img)
	bounds := newImg.Bounds()

	if bounds.Dx() != 40 || bounds.Dy() != 40 {
		t.Errorf("Expected bounds to be 40x40 (was: %dx%d)", bounds.Dx(), bounds.Dy())
	}
}

func TestResize_ZeroWidth(t *testing.T) {
	img := image.NewGray(image.Rect(0, 0, 100, 100))
	newImg := Resize(0, 40, img)
	bounds := newImg.Bounds()

	if bounds.Dx() != 40 || bounds.Dy() != 40 {
		t.Errorf("Expected bounds to be 40x40 (was: %dx%d)", bounds.Dx(), bounds.Dy())
	}
}

func TestResize_ZeroHeight(t *testing.T) {
	img := image.NewGray(image.Rect(0, 0, 100, 100))
	newImg := Resize(40, 0, img)
	bounds := newImg.Bounds()

	if bounds.Dx() != 40 || bounds.Dy() != 40 {
		t.Errorf("Expected bounds to be 40x40 (was: %dx%d)", bounds.Dx(), bounds.Dy())
	}
}

func TestResize_NoResizingRequired(t *testing.T) {
	img := image.NewGray(image.Rect(0, 0, 100, 100))
	newImg := Resize(400, 0, img)
	bounds := newImg.Bounds()

	if bounds.Dx() != 100 || bounds.Dy() != 100 {
		t.Errorf("Expected bounds to be 100x100 (was: %dx%d)", bounds.Dx(), bounds.Dy())
	}
}
