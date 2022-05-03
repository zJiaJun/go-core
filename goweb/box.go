package goweb

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}
type BoxList []Box

func (b Box) volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) setColor(c Color) {
	b.color = c
}

func (bl BoxList) biggestColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) paintItBlack() {
	for i := range bl {
		bl[i].setColor(BLACK)
		//(&bl[i]).setColor(BLACK)
	}
}

func (c Color) string() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func boxExample() {
	defer fmt.Println(common.Separator)
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}
	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.string())
	fmt.Println("The biggest one is", boxes.biggestColor().string())

	fmt.Println("Let's paint them all black")
	boxes.paintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.string())

	fmt.Println("Obviously, now, the biggest one is", boxes.biggestColor().string())

}
