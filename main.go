package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Sorting Algorithms")
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}

func NewGame() *Game {
	return &Game{}
}
