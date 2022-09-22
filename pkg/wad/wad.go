package wad

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Wad struct {
	palettes        paletteAtlas
	sprites         spriteAtlas
	tileDefinitions tileDefinitions
	backgrounds     backgrounds
}

type wadLookup struct {
	Palettes    map[string]string
	Sprites     map[string]string
	Tiles       map[string]string
	Backgrounds map[string]string
}

func New(path string) Wad {
	wadFile, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Fatal(err)
	}

	var data wadLookup
	err2 := yaml.Unmarshal(wadFile, &data)
	if err2 != nil {
		logrus.Fatal(err2)
	}

	p := paletteAtlas{}
	p.Parse(data.Palettes)

	s := spriteAtlas{}
	s.Parse(data.Sprites)

	t := tileDefinitions{}
	t.Parse(data.Tiles)

	b := backgrounds{}
	b.Parse(data.Backgrounds)

	return Wad{
		palettes:        p,
		sprites:         s,
		tileDefinitions: t,
		backgrounds:     b,
	}
}

func (w Wad) GetSprites() spriteAtlas {
	return w.sprites
}

func (w Wad) GetPalettes() paletteAtlas {
	return w.palettes
}

func (w Wad) GetTileDefinitions() tileDefinitions {
	return w.tileDefinitions
}

func (w Wad) GetBackgrounds() backgrounds {
	return w.backgrounds
}
