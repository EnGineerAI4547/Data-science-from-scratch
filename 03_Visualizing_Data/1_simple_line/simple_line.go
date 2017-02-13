package main

import (
	"fmt"
	"image/color"
	"path/filepath"
	"strings"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
	"github.com/gonum/plot/vg/draw"

	"github.com/twgophers/browser"
)

func main() {
	years := strings.Fields("1950 1960 1970 1980 1990 2000 2010")
	gdp := []float64{300.2, 543.3, 1075.9, 2862.5, 5979.6, 10289.7, 14958.3}

	xys := make(plotter.XYs, len(gdp))
	for i, amount := range gdp {
		xys[i].X = float64(i)
		xys[i].Y = amount
	}

	p, err := plot.New()
	check(err)

	p.Title.Text = "Nominal GDP"
	p.X.Label.Text = "Years"
	p.NominalX(years...)
	p.Y.Label.Text = "Billions of $"

	line, pts, err := plotter.NewLinePoints(xys)
	pts.Shape = draw.CircleGlyph{}
	pts.Color = color.RGBA{G: 128, A: 255}
	check(err)
	p.Add(line, pts)

	// Save the plot to a PNG file.
	fileName := "simple_line.png"
	err = p.Save(14*vg.Centimeter, 10*vg.Centimeter, fileName)
	check(err)

	// Display the file
	fullPath, err := filepath.Abs(fileName)
	check(err)
	url := "file://" + fullPath
	if !browser.Open(url) {
		fmt.Println("Could not open browser on: " + url)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
