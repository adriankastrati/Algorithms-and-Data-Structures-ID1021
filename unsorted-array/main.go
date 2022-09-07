package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"unsorted-array/sorted"
)

func main() {
	var (
		valueslice []int
		sliceSize  int
		//keySlice    []int
		averageTime float64
		key         int
		t0          time.Time
		min         float64
	)

	sliceSize = 5
	for mult := 1; sliceSize < 64000000; mult++ {
		averageTime = 0
		min = math.MaxFloat64

		for i := 0; i < 100; i++ {
			sliceSize = int(math.Pow(2, float64(mult)))

			valueslice = sorted.GenerateSortedList(sliceSize)

			//keySlice = sorted.GenerateSortedList(sliceSize)

			key = rand.Intn(sliceSize * 5)

			t0 = time.Now()

			if sorted.Binary_search(valueslice, key) {
				averageTime = float64(time.Since(t0))
			}

			if averageTime < min {
				min = averageTime
			}
		}
		fmt.Printf("%d %f\n", sliceSize, averageTime)

	}

}
