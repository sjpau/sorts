package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	sorting     bool
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
		self.sorting = false
	}
	l := len(self.rngvslice)
	for i := 0; i < l; i++ {
		self.matrix[i].Value = self.rngvslice[i]
		self.matrix[i].Update()
	}
	if inpututil.IsKeyJustPressed(ebiten.Key1) && self.sorting == false {
		go algorithm.Selection(self.rngvslice, 1*time.Millisecond)
		self.sorting = true
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) && self.sorting == false {
		go algorithm.Gnome(self.rngvslice, 1*time.Millisecond)
		self.sorting = true
	}
	if inpututil.IsKeyJustPressed(ebiten.Key3) && self.sorting == false {
		go algorithm.Bubble(self.rngvslice, 1*time.Millisecond)
		self.sorting = true
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
