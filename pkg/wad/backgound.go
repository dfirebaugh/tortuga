package wad

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

type backgrounds map[string]string

func (b backgrounds) Parse(backgrounds map[string]string) {
	for k, v := range backgrounds {
		f, err := ioutil.ReadFile(v)
		if err != nil {
			logrus.Fatal("unable to read file: ", v, err)
		}

		b[k] = string(f)
	}
}

func (b backgrounds) Get(key string) string {
	return b[key]
}
