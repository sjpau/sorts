package algorithm

import (
	"sort"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Selection(arr []int, delay time.Duration) {
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

	var n = len(arr)
	for !sort.IntsAreSorted(arr) {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			done <- true
			break
		}
		if i < n {
			var minIdx = i
			for j := i; j < n; j++ {
				if arr[j] < arr[minIdx] {
					minIdx = j
				}
			}
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
	done <- true
}
