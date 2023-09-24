package main

import (
	"image/color"
	"log"

	"github.com/shomali11/gridder"
)

func main() {
	imgConfig := gridder.ImageConfig{
		Name: "grid.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              20,
		Columns:           20,
		LineStrokeWidth:   2,
		BorderStrokeWidth: 10,
		LineColor:         color.Black,
		BorderColor:       color.Black,
		BackgroundColor:   color.White,
	}
	grid, err := gridder.New(imgConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}
	if err := grid.SavePNG(); err != nil {
		log.Fatal(err)
	}
}
