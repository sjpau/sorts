package algorithm

import "time"

func Cocktail(arr []int, delay time.Duration) {
	last := len(arr) - 1
	for {
		flag := false
		for i := 0; i < last; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				flag = true
			}
		}
		if !flag {
			return
		}
		flag = false
		for i := last - 1; i >= 0; i-- {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				flag = true
			}
		}
		if !flag {
			return
		}
		time.Sleep(delay)
	}
}
