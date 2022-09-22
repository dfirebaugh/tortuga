package wad

import (
	"encoding/csv"
	"image/color"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

type paletteAtlas map[string][]color.Color

func (p paletteAtlas) Parse(paletteLookups map[string]string) {
	for k, filePath := range paletteLookups {
		f, err := os.Open(filePath)
		if err != nil {
			logrus.Fatal("unable to read input file: ", filePath, err)
		}
		defer f.Close()

		csvReader := csv.NewReader(f)
		records, err := csvReader.ReadAll()
		if err != nil {
			logrus.Fatal("unable to parse palette file as csv: ", filePath, err)
		}

		var palette []color.Color
		for _, c := range records {
			r, err := strconv.Atoi(c[0])
			if err != nil {
				logrus.Fatal("unable to parse color from palette: ", filePath, err)
			}
			g, err := strconv.Atoi(c[1])
			if err != nil {
				logrus.Fatal("unable to parse color from palette: ", filePath, err)
			}
			b, err := strconv.Atoi(c[2])
			if err != nil {
				logrus.Fatal("unable to parse color from palette: ", filePath, err)
			}

			palette = append(palette, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
		}
		p[k] = palette
	}
}

func (p paletteAtlas) Get(key string) []color.Color {
	return p[key]
}
