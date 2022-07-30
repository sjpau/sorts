package algorithm

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Comb(arr []int, delay time.Duration) {
	if len(arr) < 2 {
		return
	}
	for gap := len(arr); ; {
		if gap > 1 {
			gap = gap * 4 / 5
		}
		flag := false
		for i := 0; ; {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				flag = true
				time.Sleep(delay)
				if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
					break
				}
			}
			i++
			if i+gap >= len(arr) {
				break
			}
			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				break
			}
		}
		if gap == 1 && !flag {
			time.Sleep(delay)
			break
		}
		time.Sleep(delay)
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			break
		}
	}
}
