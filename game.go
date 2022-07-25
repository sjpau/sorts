package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/sorts/algorithm"
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
	if ebiten.IsKeyPressed(ebiten.Key1) {
		self.sort = misc.Bubble
	}
	switch self.sort {
	case misc.Bubble:
		go algorithm.Bubble(self.rngvslice, 1*time.Millisecond)
		self.sort = -1
	}
	return nil
}

func (self *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	l := len(self.rngvslice)
	for i := 0; i < l; i++ {
		self.matrix[i].Draw(screen)
	}
}

func (self *Game) Layout(w, h int) (int, int) {
	f := ebiten.DeviceScaleFactor()
	return w * int(f), h * int(f)
}
