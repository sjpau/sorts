package algorithm

import (
	"sort"
	"time"
)

func Heap(list sort.Interface, delay time.Duration) {
	start := (list.Len() - 2) / 2
	end := list.Len() - 1

	for start = (list.Len() - 2) / 2; start >= 0; start-- {
		SetHeap(list, start, list.Len()-1)
		time.Sleep(delay)
	}
	for end = list.Len() - 1; end > 0; end-- {
		list.Swap(0, end)
		SetHeap(list, 0, end-1)
		time.Sleep(delay)
	}
}

func SetHeap(list sort.Interface, start, end int) {
	for root := start; root*2+1 <= end; {
		child := root*2 + 1
		if child+1 <= end && list.Less(child, child+1) {
			child++
		}
		if !list.Less(root, child) {
			return
		}
		list.Swap(root, child)
		root = child
	}
}
