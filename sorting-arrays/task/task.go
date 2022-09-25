package task

import (
	"algo/sorting-arrays/sort"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func MakeSlice(sliceSize int) []int {
	slice := make([]int, sliceSize)

	rand.Seed(time.Now().Unix())
	for i := 0; i < sliceSize; i++ {
		slice[i] = rand.Intn(sliceSize * 10)
	}

	shuffled := rand.Perm(len(slice))
	for i := 0; i < len(shuffled); i++ {
		shuffled[i] = slice[shuffled[i]]
	}

	return shuffled
}

func Task() {
	var (
		sliceSize    int
		tAverage     float64
		maxSliceSize int = 64000
	)

	for mult := 2; sliceSize < maxSliceSize; mult++ {

		tAverage = 0
		for i := 0; i < 100; i++ {
			sliceSize = int(math.Pow(2, float64(mult)))
			slice := MakeSlice(sliceSize)

			t0 := time.Now()

			sort.MergeSort(slice)

			tAverage += float64(time.Since(t0) / 100)
		}

		fmt.Printf("%d %f\n", sliceSize, tAverage/1000)
	}
}
