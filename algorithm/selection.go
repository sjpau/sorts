package algorithm

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Selection(arr []int, delay time.Duration) {
	length := len(arr)
	last := length - 1
	for i := 0; i < last; i++ {
		aMin := arr[i]
		iMin := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < aMin {
				aMin = arr[j]
				iMin = j
			}
		}
		arr[i], arr[iMin] = aMin, arr[i]
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			break
		}
		time.Sleep(delay)
	}
}
