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

func Encode(spriteNums []uint8) string {
	result := ""
	for _, c := range spriteNums {
		n := ""
		switch c {
		case 0:
			n = "0"
		case 1:
			n = "1"
		case 2:
			n = "2"
		case 3:
			n = "3"
		case 4:
			n = "4"
		case 5:
			n = "5"
		case 6:
			n = "6"
		case 7:
			n = "7"
		case 8:
			n = "8"
		case 9:
			n = "9"
		case 10:
			n = "a"
		case 11:
			n = "b"
		case 12:
			n = "c"
		case 13:
			n = "d"
		case 14:
			n = "e"
		case 15:
			n = "f"
		}

		result += n
	}

	return result
}

func EncodeCStruct(spriteNums []uint8) string {
	result := fmt.Sprintf(`
static const char sprite_pix[%d] = 
{`, len(spriteNums)/2)

	for i, c := range spriteNums {
		n := ""

		if i%16 == 0 || i == 0 {
			n += "\n\t"
		}

		if i%2 == 0 {
			n += "0x"
		}
		switch c {
		case 0:
			n += "f"
		case 1:
			n += "0"
		case 2:
			n += "1"
		case 3:
			n += "2"
		case 4:
			n += "3"
		case 5:
			n += "4"
		case 6:
			n += "5"
		case 7:
			n += "6"
		case 8:
			n += "7"
		case 9:
			n += "8"
		case 10:
			n += "9"
		case 11:
			n += "a"
		case 12:
			n += "b"
		case 13:
			n += "c"
		case 14:
			n += "d"
		case 15:
			n += "e"
		}

		if i%2 == 1 && i != len(spriteNums)-1 {
			n += ", "
		}

		result += n
	}

	result += `
};

static const char shared_palette[16][3] = 
{
	{127, 36, 84},
	{28, 43, 83},
	{0, 135, 81},
	{171, 82, 54},
	{96, 88, 79},
	{195, 195, 198},
	{255, 241, 233},
	{237, 27, 81},
	{250, 162, 27},
	{247, 236, 47},
	{93, 187, 77},
	{81, 166, 220},
	{131, 118, 156},
	{241, 118, 166},
	{252, 204, 171},
};
`

	return result
}

func Decode(spriteString string) (pixels []uint8) {
	return Parse(spriteString)
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

	if len(s.Animations) == 0 {
		return
	}
	frameIndex := int(time.Now().Unix()) % len(s.Animations[animationIndex])
	for i, pixel := range s.Animations[animationIndex][frameIndex] {
		d.Put(
			int16(x+float64(i%s.Width)),
			int16(y+float64(i/s.Width)),
			pixel)
	}
}
