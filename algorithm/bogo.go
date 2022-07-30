package algorithm

import (
	"sort"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sjpau/sorts/misc"
)

/*simple and elegant... yet efficient!*/
func Bogo(arr []int, delay time.Duration) {
	for !sort.IntsAreSorted(arr) {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			break
		}
		misc.Shuffle(arr)
		time.Sleep(delay * 20)
	}
}
