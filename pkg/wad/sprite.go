package wad

import (
	"io/ioutil"

	"github.com/dfirebaugh/tortuga/pkg/sprite"

	"github.com/sirupsen/logrus"
)

type spriteAtlas map[string]sprite.Sprite

func (s spriteAtlas) Parse(sprites map[string]string) {
	for k, v := range sprites {
		f, err := ioutil.ReadFile(v)
		if err != nil {
			logrus.Fatal("unable to read file: ", v, err)
		}

		sd := &sprite.Sprite{}
		sprite.Unmarshal(f, sd)
		s[k] = *sd
	}
}

func (s spriteAtlas) Get(key string) sprite.Sprite {
	return s[key]
}
