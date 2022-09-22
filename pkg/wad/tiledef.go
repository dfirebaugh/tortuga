package wad

import (
	"io/ioutil"
	"tortuga/pkg/sprite"

	"github.com/sirupsen/logrus"
)

type tileDefinitions map[rune][]uint8

func (t tileDefinitions) Parse(tileDefs map[string]string) {
	for k, v := range tileDefs {
		f, err := ioutil.ReadFile(v)
		if err != nil {
			logrus.Fatal("unable to read file: ", v)
		}

		t[rune(k[0])] = sprite.Parse(string(f))
	}
}

func (t tileDefinitions) Get(symbol rune) []uint8 {
	return t[symbol]
}
