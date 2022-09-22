package sprite

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type displayer interface {
	Put(x int16, y int16, c uint8)
}

type Sprite struct {
	Animations [][][]uint8
	Width      int
}

type spriteDefintion map[string][]string

func New() Sprite {
	return Sprite{}
}

func Unmarshal(b []byte, s *Sprite) {
	def := spriteDefintion{}
	yaml.Unmarshal(b, def)

	for _, animations := range def {
		frames := [][]uint8{}
		for _, animation := range animations {
			frames = append(frames, Parse(animation))
		}
		s.Animations = append(s.Animations, frames)
	}
}

func Parse(spriteString string) (pixels []uint8) {
	s := strings.ReplaceAll(strings.ReplaceAll(spriteString, "\n", ""), "\t", "")
	for _, c := range strings.Split(s, "") {
		n := 0

		switch c {
		case "0":
			n = 0
		case "1":
			n = 1
		case "2":
			n = 2
		case "3":
			n = 3
		case "4":
			n = 4
		case "5":
			n = 5
		case "6":
			n = 6
		case "7":
			n = 7
		case "8":
			n = 8
		case "9":
			n = 9
		case "a":
			n = 10
		case "b":
			n = 11
		case "c":
			n = 12
		case "d":
			n = 13
		case "e":
			n = 14
		case "f":
			n = 15
		default:
			fmt.Errorf("invalid char: %s", c)
		}

		pixels = append(pixels, uint8(n))
	}

	return pixels
}

func (s Sprite) GetPixels(animationIndex int, frameIndex int) []uint8 {
	return s.Animations[animationIndex][frameIndex]
}

func (s Sprite) VerticalMirror(animationIndex int, frameIndex int) []uint8 {
	pixels := []uint8{}

	rows := [][]uint8{}
	// frame := s.Animations[animationIndex][frameIndex]
	rowIndex := 0
	if s.Width == 0 {
		s.Width = 8
	}

	for i, p := range s.Animations[animationIndex][frameIndex] {
		if i%s.Width == 0 {
			rowIndex++
		}
		rows[rowIndex] = append(rows[rowIndex], p)
	}

	// for i := 0; i < s.Width; i++ {
	// 	pixels = append(pixels, frame[len(frame)-i])
	// }

	for _, row := range rows {
		for i := range row {
			pixels = append(pixels, row[len(row)-i])
		}
	}

	return pixels
}

func (s Sprite) DrawPixels(d displayer, pixels []uint8, x float64, y float64) {
	if s.Width == 0 {
		s.Width = 8
	}

	for i, pixel := range pixels {
		d.Put(
			int16(x+float64(i%s.Width)),
			int16(y+float64(i/s.Width)),
			pixel)
	}
}

func (s Sprite) Draw(d displayer, animationIndex int, x float64, y float64) {
	if s.Width == 0 {
		s.Width = 8
	}
	frameIndex := int(time.Now().Unix()) % len(s.Animations[animationIndex])
	for i, pixel := range s.Animations[animationIndex][frameIndex] {
		d.Put(
			int16(x+float64(i%s.Width)),
			int16(y+float64(i/s.Width)),
			pixel)
	}
}
