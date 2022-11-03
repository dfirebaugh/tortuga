package file

import (
	"io/ioutil"
	"os"

	"github.com/dfirebaugh/tortuga/pkg/sprite"
	"github.com/sirupsen/logrus"
)

func Save(frames [][]uint8) {
	output := `animation:`
	for _, frame := range frames {
		output += "\n    - " + sprite.Encode(frame)
	}
	output += "\n"

	println(output)
	os.WriteFile("./spritefile.sprite", []byte(output), 0644)
}

func Load(filePath string) [][][]uint8 {
	rawFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		logrus.Fatal(err)
	}

	s := &sprite.Sprite{}
	sprite.Unmarshal(rawFile, s)

	return s.Animations
}
