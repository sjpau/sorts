package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/sorts/composer"
	"golang.org/x/image/colornames"
)

type Game struct {
	sceneScreen *ebiten.Image
	fullscreen  bool
}

func (self *Game) Update() error {
	f := ebiten.IsKeyPressed(ebiten.KeyF)
	if !f {
		self.fullscreen = false
	} else if !self.fullscreen {
		self.fullscreen = true
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	return nil
}

func (self *Game) Draw(screen *ebiten.Image) {
	if self.sceneScreen == nil {
		self.sceneScreen = ebiten.NewImage(composer.Width, composer.Height)
	}
	screen.Fill(colornames.White)
}

func (self *Game) Layout(w, h int) (int, int) {
	f := ebiten.DeviceScaleFactor()
	return w * int(f), h * int(f)
}
