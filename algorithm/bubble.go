package algorithm

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Bubble(arr []int, delay time.Duration) {
	sorted := false
	length := len(arr)
	for !sorted {
		sorted = true
		for i := length - 1; i > 0; i-- {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				sorted = false
			}
		}
		time.Sleep(delay)
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			break
		}
	}
}
