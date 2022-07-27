package algorithm

import (
	"sort"
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
				if i >= length-1 {
					i = 0
				} else {
					i++
				}
			}
		}
	}()
	for !sort.IntsAreSorted(arr) {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			done <- true
			break
		}
		if i != length-1 && arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	done <- true
}
