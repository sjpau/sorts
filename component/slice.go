package component

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/sorts/composer"
	"golang.org/x/image/colornames"
)

type Element struct {
	this  *Object
	index int
	value int
}

type Matrix struct {
	slice []*Element
	delay time.Duration
}

func NewSlice(length int) []*Element {
	self := make([]*Element, length)
	for i := range self {
		self[i] = &Element{
			index: i,
			this: &Object{
				X:      i,
				Y:      rand.Intn(composer.Height), // here might be a hiccup
				Width:  1,
				Height: 1,
				Image:  ebiten.NewImage(1, 1),
			},
		}
		self[i].this.Image.Fill(colornames.White)
	}
	return self
}

func (self *Element) Draw(screen *ebiten.Image) {
	o := &ebiten.DrawImageOptions{}
	o.GeoM.Scale(1, 1)
	o.GeoM.Translate(float64(self.this.X), float64(self.this.Y))
	screen.DrawImage(self.this.Image, o)
}

func NewMatrix(delay time.Duration) *Matrix {
	s := NewSlice(composer.Width)
	self := Matrix{
		slice: s,
		delay: delay,
	}
	return &self
}

func (self *Matrix) Draw(screen *ebiten.Image) {
	if self.slice != nil {
		for i := range self.slice {
			self.slice[i].Draw(screen)
		}
	}
}
