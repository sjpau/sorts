package algorithm

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Gnome(arr []int, delay time.Duration) {
	length := len(arr)
	for i, j := 1, 2; i < length; {
		if arr[i-1] > arr[i] {
			arr[i-1], arr[i] = arr[i], arr[i-1]
			i--
			if i > 0 {
				continue
			}
		}
		i = j
		j++
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			break
		}
		time.Sleep(delay)
	}
}
