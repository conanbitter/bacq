package main

func main() {
	iimg, err := imageLoad("tests/test05lo.png")
	if err != nil {
		panic(err)
	}
	data := imageToData(iimg)
	cq := NewQuantizier(16, 1000, 5)
	cq.Input(data)
	cq.Run()
	pal := cq.GetPalette()
	convertImage(iimg, "tests/test05lo_conv16.png", pal, IndexerPattern4)
}
