package main

import (
	"errors"
	"fmt"
	"image"
	"log"
	"os"

	"image/color"
	_ "image/jpeg"
	"image/png"

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

func brightnessCorrection(index int, level int, totalLevels int) int {
	if level == totalLevels {
		return index
	}
	if level == 0 {
		return -1
	}
	currentLevel := index%totalLevels + 1
	color := index / totalLevels
	newLevel := int(float64(currentLevel) * float64(level) / float64(totalLevels))
	if newLevel == 0 {
		return -1
	}
	return color*totalLevels + newLevel - 1
}

func dataToImageF(data []FloatColor, outputFilename string, width int, height int) {
	oimg := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			index := y*width + x
			col := data[index].ToIntColor()
			oimg.SetRGBA(x, y, color.RGBA{uint8(col.R), uint8(col.G), uint8(col.B), 255})
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

func convertImage(inputImage any, outputFilename string, palette any, indexer ImageIndexer, level uint, totalLevels uint) {
	var err error

	var pal Palette
	switch palt := palette.(type) {
	case Palette:
		pal = palt
	case string:
		pal = PaletteLoad(palt)
	default:
		panic(errors.New("wrong palette type"))
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
			colp := brightnessCorrection(imgIndexed[index], int(level), int(totalLevels))
			if colp >= pal.Len() {
				panic(fmt.Errorf("colors: %d, level:%d, original:%d", colp, level, imgIndexed[index]))
			}
			if colp == -1 {
				colp = pal.Len() - 1
			}
			if colp < -1 {
				panic(fmt.Errorf("colors: %d, level:%d, original:%d", colp, level, imgIndexed[index]))
			}
			coli := pal[colp]
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
