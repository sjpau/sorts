package component

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Element struct {
	Body  Object
	Value int
	Index int
}

func NewMatrix(length int) []Element {
	self := make([]Element, length)
	for i := 0; i < length; i++ {
		tmp := Element{
			Body: Object{},
		}
		tmp.Index = i
		tmp.Body.X = i
		tmp.Body.Image = ebiten.NewImage(1, 1)
		tmp.Body.Image.Fill(colornames.White)
		self[i] = tmp
	}
	return self
}

func (self *Element) Update() {
	self.Body.Y = self.Value
}

func (self *Element) Draw(screen *ebiten.Image, marginX, marginY float64) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, 1)
	o.GeoM.Translate(marginX+float64(self.Body.X), marginY+float64(self.Body.Y))
	screen.DrawImage(self.Body.Image, o)
}
