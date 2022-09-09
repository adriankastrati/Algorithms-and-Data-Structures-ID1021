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

func LadderSearch(keySlice []int, arrayValue []int) float64 {
	var (
		keyIndex   int = 0
		valueIndex int = 0
		t0         time.Time
		Dup        int
	)

	t0 = time.Now()
	for keyIndex < len(keySlice) && valueIndex < len(arrayValue) {

		if arrayValue[valueIndex] == keySlice[keyIndex] {
			Dup++
			valueIndex++
			keyIndex++
			continue
		}

		if arrayValue[valueIndex] < keySlice[keyIndex] {
			valueIndex++
			continue
		}

		if arrayValue[valueIndex] > keySlice[keyIndex] {
			keyIndex++
			continue
		}

	}

	return float64(time.Since(t0))

}
