package main

import "fmt"

func main() {
	iimg, err := imageLoad("tests/test05lo.png")
	if err != nil {
		panic(err)
	}
	data := imageToData(iimg)
	cq := NewQuantizier(16, 5, 1000, 5)
	cq.Input(data)
	cq.Run()
	pal := cq.GetPalette()
	fmt.Printf("Colors in palette: %d\n", pal.Len())
	pal.SavePreview(5, "tests/test05lo_pal.png")
	convertImage(iimg, "tests/test05lo_conv16.png", pal, IndexerPattern4)
}
