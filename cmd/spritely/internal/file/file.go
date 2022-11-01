package file

import (
	"os"

	"github.com/dfirebaugh/tortuga/pkg/sprite"
)

func Save(frames [][]uint8) {
	output := `animation:`
	for _, frame := range frames {
		output += "\n\t- " + sprite.Encode(frame)
	}
	output += "\n"

	println(output)
	os.WriteFile("./spritefile.sprite", []byte(output), 0644)
}

func Load(filePath string) {
	// f, err := os.Open(filePath)
	// if err != nil {
	// 	logrus.Error(err)
	// }

	// println(f)
}
