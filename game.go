package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Game struct {
	fullscreen bool
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
	screen.Fill(colornames.White)
}

func (self *Game) Layout(w, h int) (int, int) {
	f := ebiten.DeviceScaleFactor()
	return w * int(f), h * int(f)
}
