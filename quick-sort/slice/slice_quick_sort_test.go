package sliceQuickSort

import (
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	slice := make([]int, 10)

	for i, _ := range slice {
		slice[i] = rand.Intn(35)
	}
	QuickSort(slice, 2, 6)

	for i := 2; i < 6; i++ {
		if slice[i] > slice[i+1] {
			t.Errorf("not sorted correctly: %d bigger than %d", slice[i], slice[i+1])
		}
	}
}
