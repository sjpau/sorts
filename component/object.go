package component

import "github.com/hajimehoshi/ebiten/v2"

type Object struct {
	Image  *ebiten.Image
	X      int
	Y      int
	Width  int
	Height int
}
