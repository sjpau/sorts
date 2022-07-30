package algorithm

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Bubble(arr []int, delay time.Duration) {
	ticker := time.NewTicker(delay)
	done := make(chan bool)
	i := 0
	length := len(arr)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				if i > 0 {
					i--
				} else {
				}
			}
		}
	}()

	sorted := false
	for !sorted {
		sorted = true
		for i = length - 1; i > 0; {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				sorted = false
			}

			if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
				done <- true
				break
			}
		}
	}
	done <- true
}
