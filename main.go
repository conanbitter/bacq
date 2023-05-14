package main

import "fmt"

func main() {
	/*levels := []struct {
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
	}*/

	/*levels2 := []struct {
		levels uint
		colors int
	}{
		{5, 51},
		{8, 31},
		{10, 25},
		{16, 15},
	}*/

	iimg, err := imageLoad("tests/test01lo.png")
	if err != nil {
		panic(err)
	}
	data := imageToData(iimg)
	width := iimg.Bounds().Size().X
	height := iimg.Bounds().Size().Y
	/*for _, lvl := range levels2 {
		fmt.Printf("%d x %d\n", lvl.levels, lvl.colors)
		cq := NewQuantizier(lvl.colors, lvl.levels, 1000, 5)
		cq.Input(data)
		cq.Run()
		pal := cq.GetPalette()
		fmt.Printf("Colors in palette: %d\n", pal.Len())
		pal.SavePreview(int(lvl.levels), fmt.Sprintf("tests/test01lo_pal%d.png", lvl.levels))
		for l := int(lvl.levels); l >= 0; l-- {
			convertImage(
				iimg,
				fmt.Sprintf("tests/test01lo_lvl%dx%d_%d.png", lvl.levels, lvl.colors, l),
				pal,
				IndexerFS,
				uint(l),
				lvl.levels)
		}
	}*/
	for i := 10; i >= 0; i-- {
		k := float64(i) / 10.0
		ldata := make([]FloatColor, width*height)
		for j, c := range data {
			col := c.ToFloatColor()
			col.R *= k
			col.G *= k
			col.B *= k
			ldata[j] = col
		}
		dataToImageF(ldata, fmt.Sprintf("tests/test01lo_fc_%d.png", 10-i), width, height)
	}
}
