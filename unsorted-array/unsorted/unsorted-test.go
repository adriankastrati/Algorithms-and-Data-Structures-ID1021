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

func ShuffleSlice(slice []int) []int {
	shuffled := rand.Perm(len(slice))

	for i := 0; i < len(shuffled); i++ {
		shuffled[i] = slice[shuffled[i]]
	}

	return shuffled
}

func MakeSlice(sliceSize int) []int {
	slice := make([]int, sliceSize)

	for i := 0; i < sliceSize; i++ {
		slice[i] = rand.Intn(sliceSize * 10)
	}

	return slice
}
