package main

import (
	"image"
	"os"

	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/tiff"
)

func imageLoad(filename string) (image.Image, error) {
	imgFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func imageToData(img image.Image) []IntColor {
	bounds := img.Bounds()
	result := make([]IntColor, 0, (bounds.Max.X-bounds.Min.X)*(bounds.Max.Y-bounds.Min.Y))
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			result = append(result, IntColor{int(r / 257), int(g / 257), int(b / 257)}.Normalized())
		}
	}
	return result
}
