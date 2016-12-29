package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
)

/*

Code from https://github.com/SimonWaldherr/cgolGo/

*/
type Field struct {
	cells  [][]int
	width  int
	height int
	gv     *GifVisualizer
}

func (field *Field) setTrace(loc string) {
	field.gv = &GifVisualizer{}
	field.gv.Setup(loc)
}

func newField(width, height int) *Field {
	cells := make([][]int, height)
	for cols := range cells {
		cells[cols] = make([]int, width)
	}
	return &Field{cells: cells, width: width, height: height}
}

func (field *Field) setVitality(x, y int, vitality int) {
	x += field.width
	x %= field.width
	y += field.height
	y %= field.height
	field.cells[y][x] = vitality
}

func (field *Field) getVitality(x, y int) int {
	x += field.width
	x %= field.width
	y += field.height
	y %= field.height
	return field.cells[y][x]
}

func (field *Field) GetVitality(x, y int) int {
	return field.getVitality(x, y)
}

func (field *Field) nextVitality(x, y int) int {
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && (field.getVitality(x+i, y+j) > 0) {
				alive++
			}
		}
	}
	vitality := field.getVitality(x, y)
	if alive == 3 || alive == 2 && (vitality > 0) {
		if vitality < 8 {
			return vitality + 1
		} else {
			return vitality
		}
	}
	return 0
}

func generateFirstRound(width, height int) *Field {
	field := newField(width, height)
	for i := 0; i < (width * height / 4); i++ {
		field.setVitality(rand.Intn(width), rand.Intn(height), 1)
	}
	return field
}

func loadFromBitwiseRepresention(width, height, offsetw, offseth, bpw, bph int, brp *big.Int) *Field {
	field := newField(width, height)
	count := 0
	for cw := 1; cw <= bpw; cw++ {
		for ch := 1; ch <= bph; ch++ {
			b := int(brp.Bit(count))
			field.setVitality(cw+offsetw, ch+offseth, b)
			count++
		}
	}
	return field
}
func (field *Field) Finialize() {
	if field.gv != nil {
		field.gv.Complete()
	}
}

func (field *Field) CountLife() int {
	count := 0
	for cw := 1; cw <= field.width; cw++ {
		for ch := 1; ch <= field.height; ch++ {
			if field.getVitality(ch, cw) != 0 {
				count++
			}
		}
	}
	return count
}

func loadFirstRound(width, height int, filename string) *Field {
	finfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println(filename + " doesn't exist")
		return generateFirstRound(width, height)
	} else {
		if finfo.IsDir() {
			fmt.Println(filename + " is a directory")
			return generateFirstRound(width, height)
		} else {
			field := newField(width, height)
			gofile, _ := ioutil.ReadFile(filename)
			output := []rune(string(gofile))
			x := 0
			y := 0
			for _, char := range output {
				switch char {
				case 10:
					y++
					x = 0
				case 49:
					field.setVitality(x, y, 1)
				case 50:
					field.setVitality(x, y, 2)
				case 51:
					field.setVitality(x, y, 3)
				case 52:
					field.setVitality(x, y, 4)
				case 53:
					field.setVitality(x, y, 5)
				case 54:
					field.setVitality(x, y, 6)
				case 55:
					field.setVitality(x, y, 7)
				case 56:
					field.setVitality(x, y, 8)
				case 57:
					field.setVitality(x, y, 9)
				default:
					if char != 32 {
						field.setVitality(x, y, 1)
					} else {
						field.setVitality(x, y, 0)
					}
				}
				x++
			}
			return field
		}
	}
	return generateFirstRound(width, height)
}

func (field *Field) nextRound() *Field {
	if field.gv != nil {
		field.gv.AddFrame(field.cells)
	}
	new_field := newField(field.width, field.height)
	for y := 0; y < field.height; y++ {
		for x := 0; x < field.width; x++ {
			new_field.setVitality(x, y, field.nextVitality(x, y))
		}
	}
	field.cells = new_field.cells
	return new_field
}

func (field *Field) printField() string {
	var buffer bytes.Buffer
	for y := 0; y < field.height; y++ {
		for x := 0; x < field.width; x++ {
			if field.getVitality(x, y) > 0 {
				buffer.WriteString("█")
			} else {
				buffer.WriteByte(byte(' '))
			}
		}
		buffer.WriteByte('\n')
	}
	return buffer.String()
}

type GifVisualizer struct {
	name string
	g    *gif.GIF
}

func (gv *GifVisualizer) Setup(name string) {
	gv.g = &gif.GIF{
		LoopCount: 1,
	}
	gv.name = name
}

func (gv *GifVisualizer) AddFrame(arr [][]int) {
	frame := buildImage(arr)
	gv.g.Image = append(gv.g.Image, frame)
	gv.g.Delay = append(gv.g.Delay, 40)
}

func (gv *GifVisualizer) Complete() {
	writeGif(gv.name, gv.g)
}

func buildImage(arr [][]int) *image.Paletted {
	var frame = image.NewPaletted(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{len(arr[0]), len(arr)},
		},
		color.Palette{
			color.Gray{uint8(255)},
			color.Gray{uint8(0)},
		},
	)

	for x, xv := range arr {
		for y, yv := range xv {
			if yv > 0 {
				frame.SetColorIndex(y, x, uint8(1))
			}
		}
	}
	return frame
}

func writeGif(name string, g *gif.GIF) {
	w, err := os.Create(name + ".gif")
	if err != nil {
		fmt.Println("os.Create")
		panic(err)
	}
	err = gif.EncodeAll(w, g)
	if err != nil {
		fmt.Println("gif.EncodeAll")
		panic(err)
	}
	if err := w.Close(); err != nil {
		panic(err)
	}
}
