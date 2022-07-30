package algorithm

import (
	"sort"
	"time"

	"github.com/sjpau/sorts/misc"
)

/*simple and elegant... yet efficient!*/
func Bogo(arr []int, delay time.Duration) {

	for !sort.IntsAreSorted(arr) {
		misc.Shuffle(arr)
		time.Sleep(delay * 25)
	}
}
