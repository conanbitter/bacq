package main

import "fmt"

func main() {
	levels := []struct {
		levels uint
		colors int
	}{
		{2, 127},
		{3, 85},
		{4, 63},
		{5, 51},
		{6, 42},
		{7, 36},
		{8, 31},
		{9, 28},
		{10, 25},
		{11, 23},
		{12, 21},
		{13, 19},
		{14, 18},
		{15, 17},
		{17, 15},
		{18, 14},
		{19, 13},
		{21, 12},
		{23, 11},
		{25, 10},
		{28, 9},
		{31, 8},
		{36, 7},
		{42, 6},
		{51, 5},
		{63, 4},
		{85, 3},
		{127, 2},
		{255, 1},
	}

	iimg, err := imageLoad("tests/test02lo.png")
	if err != nil {
		panic(err)
	}
	data := imageToData(iimg)
	for _, lvl := range levels {
		fmt.Printf("%d x %d\n", lvl.levels, lvl.colors)
		cq := NewQuantizier(lvl.colors, lvl.levels, 1000, 5)
		cq.Input(data)
		cq.Run()
		pal := cq.GetPalette()
		fmt.Printf("Colors in palette: %d\n", pal.Len())
		//pal.SavePreview(int(lvl.levels), fmt.Sprintf("tests/test07lo_pal%d.png", lvl.levels))
		convertImage(iimg, fmt.Sprintf("tests/test02lo_bacq%dx%d.png", lvl.levels, lvl.colors), pal, IndexerPosterize)
	}
}
