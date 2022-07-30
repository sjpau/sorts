package algorithm

import "time"

func Quick(arr []int, delay time.Duration) {
	var pex func(int, int)
	pex = func(lower, upper int) {
		for {
			time.Sleep(delay)
			switch upper - lower {
			case -1, 0:
				return
			case 1:
				if arr[upper] < arr[lower] {
					arr[upper], arr[lower] = arr[lower], arr[upper]
				}
				return
			}
			bx := (upper + lower) / 2
			b := arr[bx]
			lp := lower
			up := upper
		outer:
			for {
				time.Sleep(delay)
				for lp < upper && !(b < arr[lp]) {
					lp++
				}
				for {
					if lp > up {
						break outer
					}
					if arr[up] < b {
						break
					}
					up--
				}
				arr[lp], arr[up] = arr[up], arr[lp]
				lp++
				up--
			}
			if bx < lp {
				if bx < lp-1 {
					arr[bx], arr[lp-1] = arr[lp-1], b
				}
				up = lp - 2
			} else {
				if bx > lp {
					arr[bx], arr[lp] = arr[lp], b
				}
				up = lp - 1
				lp++
			}
			if up-lower < upper-lp {
				pex(lower, up)
				lower = lp
			} else {
				pex(lp, upper)
				upper = up
			}
		}
	}
	pex(0, len(arr)-1)
}
