package main

import (
	"fmt"
	"github.com/h2non/bimg"
	"os"
)

func main() {
	buffer, err := bimg.Read("test.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	watermark := bimg.Watermark{
		Text:       "By EasyPDF",
		Opacity:    0.25,
		Width:      200,
		DPI:        100,
		Margin:     150,
		Font:       "sans bold 12",
		Background: bimg.Color{255, 255, 255},
	}

	newImage, err := bimg.NewImage(buffer).Watermark(watermark)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	bimg.Write("new.jpg", newImage)
}
