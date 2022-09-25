package introduction

import (
	"fmt"
	"math/rand"
	"time"
)

func random_array_access(loops int, n int) {
	var (
		totalTime    time.Duration
		totalAddTime time.Duration
		mean         float64
		sum          int = 0
		slice            = make([]int, n)
		sliceIndex       = make([]int, n)
	)

	for i := 0; i < n; i++ {
		sliceIndex[i] = rand.Intn(n)
		slice[i] = 1
	}

	time0 := time.Now()

	for i := 0; i < loops; i++ {
		for j := 0; j < n; j++ {
			sum += slice[sliceIndex[j]]
		}
	}

	totalTime = time.Now().Sub(time0)

	time0 = time.Now()

	for i := 0; i < loops; i++ {
		for j := 0; j < n; j++ {
			sum += slice[j]
		}
	}

	totalAddTime = time.Now().Sub(time0)

	mean = float64(totalTime-totalAddTime) / float64(loops*n)
	fmt.Printf("mean resolution: %g ns\n", mean)
}

/*
func main() {
	var smallSliceSize, mediumSliceSize, largeSliceSize, largestSliceSize, loops = 100, 1000, 10000, 15000, 1000000
	fmt.Printf("Small size array: ")
	random_array_access(loops, smallSliceSize)

	fmt.Printf("Medium size array: ")
	random_array_access(loops, mediumSliceSize)

	fmt.Printf("Large size array: ")
	random_array_access(loops, largeSliceSize)

	fmt.Printf("Largest size array: ")
	random_array_access(loops, largestSliceSize)

}*/
