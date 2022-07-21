package misc

import "math/rand"

func Shuffle(a []int) []int {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
}

func RandomUniqueSlice(length int) []int {
	a := make([]int, length)
	for i := range a {
		a[i] = i
	}
	shuffled := Shuffle(a)
	return shuffled
}
