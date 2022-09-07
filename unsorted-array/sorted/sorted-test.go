package sorted

import (
	"math/rand"
	"time"
)

func SearchSortedList(slice []int, key int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == key {
			return true
		}
	}
	return false
}

func GenerateSortedList(sliceSize int) []int {
	rand.Seed(time.Now().UTC().UnixNano())

	slice := make([]int, sliceSize)

	next := 0

	for i := 0; i < sliceSize; i++ {
		next += rand.Intn(10) + 1
		slice[i] = next
	}

	return slice
}
