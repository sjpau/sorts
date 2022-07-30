package main

import (
	"sort"
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
	fullscreen  bool
}

func (self *Game) Update() error {
	fullscreen := ebiten.IsKeyPressed(ebiten.KeyF)
	if !fullscreen {
		self.fullscreen = false
	} else if !self.fullscreen {
		self.fullscreen = true
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
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
		go algorithm.Selection(self.rngvslice, 5*time.Millisecond)
		self.sorting = true
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) && self.sorting == false {
		go algorithm.Gnome(self.rngvslice, 5*time.Millisecond)
		self.sorting = true
	}
	if inpututil.IsKeyJustPressed(ebiten.Key3) && self.sorting == false {
		go algorithm.Bubble(self.rngvslice, 5*time.Millisecond)
		self.sorting = true
	}
	if inpututil.IsKeyJustPressed(ebiten.Key4) && self.sorting == false {
		go algorithm.Heap(sort.IntSlice(self.rngvslice), 5*time.Millisecond)
		self.sorting = true
	}
	if inpututil.IsKeyJustPressed(ebiten.Key5) && self.sorting == false {
		go algorithm.Cocktail(self.rngvslice, 5*time.Millisecond)
		self.sorting = true
	}
	if inpututil.IsKeyJustPressed(ebiten.Key6) && self.sorting == false {
		go algorithm.Comb(self.rngvslice, 5*time.Millisecond)
		self.sorting = true
	}
	if inpututil.IsKeyJustPressed(ebiten.Key7) && self.sorting == false {
		go algorithm.Quick(self.rngvslice, 5*time.Millisecond)
		self.sorting = true
	}
	if inpututil.IsKeyJustPressed(ebiten.Key8) && self.sorting == false {
		go algorithm.Shell(self.rngvslice, 5*time.Millisecond)
		self.sorting = true
	}
	return nil
}

func (self *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Black)
	l := len(self.rngvslice)
	marginX := 0
	marginY := 0
	if !self.fullscreen {
		w, h := screen.Size()
		marginX = (w - composer.Width) / 2
		marginY = (h - composer.Height) / 2
	}
	for i := 0; i < l; i++ {
		self.matrix[i].Draw(screen, float64(marginX), float64(marginY))
	}
}

func (self *Game) Layout(w, h int) (int, int) {
	f := ebiten.DeviceScaleFactor()
	return w * int(f), h * int(f)
}
