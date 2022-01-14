package compressor

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
)

func CompressAndExport(quality int, writer io.Writer, img image.Image) (int, error) {
	var buf bytes.Buffer
	err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	if err != nil {
		return -1, err
	}
	bytes := buf.Len()
	writer.Write(buf.Bytes())
	return bytes, nil
}
