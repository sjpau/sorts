package algorithm

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Shell(arr []int, delay time.Duration) {
	for tag := len(arr) / 2; tag > 0; tag = (tag + 1) * 5 / 11 {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			break
		}
		time.Sleep(delay)
		for i := tag; i < len(arr); i++ {
			j, tmp := i, arr[i]
			for ; j >= tag && arr[j-tag] > tmp; j -= tag {
				arr[j] = arr[j-tag]
				time.Sleep(delay)
			}
			arr[j] = tmp
		}
	}
}
