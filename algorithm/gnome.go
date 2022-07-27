package algorithm

import (
	"sort"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Gnome(arr []int, delay time.Duration) {
	ticker := time.NewTicker(delay)
	done := make(chan bool)
	i := 0
	counter := 0
	target := 0
	length := len(arr)
	flag := true
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
		flag = true
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			done <- true
			break
		}
		if i < length-1 && target < length-1 && arr[i] > arr[i+1] {
			arr[target], arr[i+1] = arr[i+1], arr[target]
			flag = false
		}
		if flag && i < length-1-counter {
			i++
			target++
		} else {
			target = 0
			i = 0
			counter++
		}
		if counter == length-1 {
			counter = 0
			target = 0
			i = 0
		}
	}
	done <- true
}
