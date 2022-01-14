package cli

import (
	"flag"
	"image/jpeg"
)

type CliArguments struct {
	InputFiles []string
	MaxWidth   uint
	Estimate   bool
	Quality    int
}

func ParseArguments() CliArguments {
	estimate := flag.Bool("estimate", false, "Only estimate new image sizes without writing them.")
	maxWidth := flag.Uint("maxWidth", 0, "Sets the maximum width for an image. The height is set accordingly.")
	quality := flag.Int("quality", jpeg.DefaultQuality, "Sets the quality (1-100).")

	flag.Parse()

	return CliArguments{
		InputFiles: flag.Args(),
		Estimate:   *estimate,
		MaxWidth:   *maxWidth,
		Quality:    *quality,
	}
}
