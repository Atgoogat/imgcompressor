package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"math"
	"os"
	"path/filepath"

	"github.com/atgoogat/imgcompressor/cli"
	"github.com/atgoogat/imgcompressor/compressor"
)

func prefixFilename(filename string) string {
	dir, file := filepath.Split(filename)
	return filepath.Join(dir, "cmp_"+file)
}

func main() {
	args := cli.ParseArguments()

	if args.MaxWidth == 0 {
		args.MaxWidth = math.MaxUint
	}
	if args.Quality < 1 || args.Quality > 100 {
		args.Quality = jpeg.DefaultQuality
	}

	bytesSum := 0

	for _, filename := range args.InputFiles {
		imgFile, err := os.Open(filename)
		if err != nil {
			err = fmt.Errorf(filename, err, "\n")
			fmt.Fprint(os.Stderr, err)
			continue
		}
		defer imgFile.Close()

		img, _, err := image.Decode(imgFile)
		if err != nil {
			err = fmt.Errorf(filename, err, "\n")
			fmt.Fprint(os.Stderr, err)
			continue
		}

		var writer io.Writer
		if args.Estimate {
			writer = &bytes.Buffer{}
		} else {
			fileHandle, err := os.Create(prefixFilename(filename))
			if err != nil {
				panic(err)
			}
			defer fileHandle.Close()
			writer = fileHandle
		}

		img = compressor.Resize(args.MaxWidth, img)

		bytes, err := compressor.CompressAndExport(args.Quality, writer, img)
		if err != nil {
			err = fmt.Errorf(filename, err, "\n")
			fmt.Fprint(os.Stderr, err)
			continue
		}
		fmt.Printf("%15s %5d kb\n", filename, bytes/1024)
		bytesSum += bytes
	}
	fmt.Printf("%15s %5d kb\n", "Sum", bytesSum/1024)
}
