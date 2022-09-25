package Task

import (
	"algo/unsorted-array/sorted"
	"fmt"
	"math"
)

func Task() {
	var (
		valueslice  []int
		sliceSize   int
		keySlice    []int
		averageTime float64
	)

	for mult := 1; sliceSize < 64000000; mult++ {
		averageTime = 0

		for i := 0; i < 100; i++ {
			sliceSize = int(math.Pow(2, float64(mult)))
			valueslice = sorted.GenerateSortedList(sliceSize)
			keySlice = sorted.GenerateSortedList(sliceSize * 5)

			averageTime += sorted.LadderSearch(keySlice, valueslice)

		}
		fmt.Printf("%d %f\n", sliceSize, averageTime/100000)

	}

}
