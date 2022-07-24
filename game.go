package main

import (
	"image/color"
	"math"

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
	fullscreen  bool
	sort        misc.SortStatus
}

func (self *Game) Update() error {
	f := ebiten.IsKeyPressed(ebiten.KeyF)
	if !f {
		self.fullscreen = false
	} else if !self.fullscreen {
		self.fullscreen = true
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}
	if self.matrix == nil {
		self.matrix = component.NewMatrix(5)
	}
	if self.rngvslice == nil {
		self.rngvslice = misc.RandomUniqueSlice(5)
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
	if self.sceneScreen == nil {
		self.sceneScreen = ebiten.NewImage(composer.Width, composer.Height)
		self.sceneScreen.Fill(color.RGBA{24, 24, 24, 255})
	}

	w, h := screen.Size()
	xscale, yscale := float64(w)/composer.Width, float64(h)/composer.Height
	scaling := int(math.Min(xscale, yscale))
	Width, Height := composer.Width*scaling, composer.Height*scaling
	ysf, xsf := float64(scaling), float64(scaling)
	marginHor := (w - Width) / 2
	marginVer := (h - Height) / 2

	o := &ebiten.DrawImageOptions{}
	if float64(scaling) < 1 {
		// Resolution is too low
		marginHor, marginVer = 0, 0
		xsf, ysf = float64(w)/composer.Width, float64(h)/composer.Height
		if w >= composer.Width {
			xsf = 1.0
			marginHor = (w - composer.Width) / 2
		}
		if h >= composer.Height {
			ysf = 1.0
			marginVer = (h - composer.Height) / 2
		}
	}
	o.GeoM.Scale(xsf, ysf)
	if marginHor != 0 || marginVer != 0 {
		screen.Fill(color.RGBA{24, 24, 24, 255})
		o.GeoM.Translate(float64(marginHor), float64(marginVer))
	}
	if scaling >= 1 {
	}
	screen.Fill(colornames.Grey)
	screen.DrawImage(self.sceneScreen, o)
	l := len(self.rngvslice)
	for i := 0; i < l; i++ {
		self.matrix[i].Draw(self.sceneScreen)
	}
}

func (self *Game) Layout(w, h int) (int, int) {
	f := ebiten.DeviceScaleFactor()
	return w * int(f), h * int(f)
}
