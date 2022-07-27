package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sjpau/sorts/composer"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ebiten.SetWindowResizable(false)
	ebiten.SetWindowTitle("Sorting Algorithms")
	ebiten.SetWindowSize(composer.Width, composer.Height)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

func NewGame() *Game {
	self := &Game{}
	self.sorting = false
	return self
}
