package algorithm

import (
	"fmt"
	"sort"
	"time"
)

func Bubble(arr []int, delay time.Duration) {
	ticker := time.NewTicker(delay)
	done := make(chan bool)
	i := 0
	j := 0
	length := len(arr)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				if i > length-1 {
					i = 0
				} else {
					i++
					fmt.Println("i", i) //800
				}
				if j == length-i-1 {
					j = 0
				} else {
					j++
					fmt.Println("j", j) //799
				}
			}
		}
	}()
	for !sort.IntsAreSorted(arr) {
		if j != length-1 && arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	done <- true
}
