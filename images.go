package main

import (
	"errors"
	"image"
	"log"
	"os"

	"image/color"
	_ "image/jpeg"
	"image/png"
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

func convertImage(inputImage any, outputFilename string, palette any, indexer ImageIndexer) {
	var err error

	var pal Palette
	switch palt := palette.(type) {
	case Palette:
		pal = palt
	case string:
		pal = PaletteLoad(palt)
	default:
		panic(errors.New("frong palette type"))
	}

	var img image.Image
	switch imgt := inputImage.(type) {
	case image.Image:
		img = imgt
	case string:
		img, err = imageLoad(imgt)
		if err != nil {
			panic(err)
		}
	}

	width := img.Bounds().Size().X
	height := img.Bounds().Size().Y

	imgData := imageToData(img)
	imgIndexed := indexer(imgData, pal, width, height)

	oimg := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			index := y*width + x
			coli := pal[imgIndexed[index]]
			oimg.SetRGBA(x, y, color.RGBA{uint8(coli.R), uint8(coli.G), uint8(coli.B), 255})
		}
	}
	outf, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer outf.Close()
	if err = png.Encode(outf, oimg); err != nil {
		log.Printf("failed to encode: %v", err)
	}
}
