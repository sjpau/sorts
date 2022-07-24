package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/sorts/component"
	"github.com/sjpau/sorts/composer"
	"github.com/sjpau/sorts/misc"
	"golang.org/x/image/colornames"
)

type Game struct {
	sceneScreen *ebiten.Image
	matrix      []component.Element
	rngvslice   []int
	sort        misc.SortStatus
}

func (self *Game) Update() error {
	if self.matrix == nil {
		self.matrix = component.NewMatrix(composer.Width)
	}
	if self.rngvslice == nil {
		self.rngvslice = misc.RandomUniqueSlice(composer.Width)
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		misc.Shuffle(self.rngvslice)
	}
	l := len(self.rngvslice)
	for i := 0; i < l; i++ {
		self.matrix[i].Value = self.rngvslice[i]
		self.matrix[i].Update()
	}
	return nil
}

func (self *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	l := len(self.rngvslice)
	for i := 0; i < l; i++ {
		self.matrix[i].Draw(screen)
		fmt.Println(self.matrix[i].Body.Y)
	}
}

func (self *Game) Layout(w, h int) (int, int) {
	f := ebiten.DeviceScaleFactor()
	return w * int(f), h * int(f)
}
