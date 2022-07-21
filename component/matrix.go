package component

import (
	"container/list"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/sorts/misc"
	"golang.org/x/image/colornames"
)

type Element struct {
	Body  Object
	Value int
}

func NewList(length int) *list.List {
	self := list.New()
	rngv := misc.RandomUniqueSlice(length)
	for i := 0; i < length; i++ {
		tmp := Element{
			Body: Object{},
		}
		tmp.Value = rngv[i]
		tmp.Body.X = i
		tmp.Body.Image = ebiten.NewImage(1, 1)
		tmp.Body.Image.Fill(colornames.White)
		self.PushBack(tmp)
	}
	return self
}

func (self *Element) Update() {
	self.Body.Y = self.Value
}

func (self *Element) Draw(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, 1)
	o.GeoM.Translate(float64(self.Body.X), float64(self.Body.Y))
	screen.DrawImage(self.Body.Image, o)
}
