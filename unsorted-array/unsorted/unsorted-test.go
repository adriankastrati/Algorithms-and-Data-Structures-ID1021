package unsorted

import (
	"math/rand"
)

func Search_unsorted(slice []int, key int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == key {
			return true
		}
	}
	return false

}

func MakeSlice(sliceSize int) []int {
	slice := make([]int, sliceSize)

	for i := 0; i < sliceSize; i++ {
		slice[i] = rand.Intn(sliceSize * 10)
	}

	return slice
}

/*
func main() {

	for i := 10; i < 10000000; i = i * 2 {

		slice := unsorted.MakeSlice(i)

		key := rand.Intn(i)
		t0 := time.Now()

		for loop := 1; loop < 1000; loop++ {

			if unsorted.Search_unsorted(slice, key) {
				break
			}

			key = rand.Intn(10 * i)
		}

		tdelta := float64(time.Since(t0))

		fmt.Printf("%d %f\n", len(slice), tdelta/1000)

	}
}

*/
