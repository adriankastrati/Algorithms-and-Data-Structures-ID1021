package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsorted-array/sorted"
	"unsorted-array/unsorted"
)

func main() {
	var (
		bin    float64 = 0
		sort   float64 = 0
		unsort float64 = 0
	)

	sliceSize := 32000000
	slice := sorted.Sorted(sliceSize)
	key := slice[rand.Intn(sliceSize)]

	loopsize := 320
	for i := 0; i < loopsize; i++ {

		for i := 0; i < sliceSize; i++ {

			t0 := time.Now()
			if sorted.Binary_search(slice, key) {
				bin += float64(time.Since(t0))
			}

			t0 = time.Now()
			for i := 0; i < len(slice); i++ {
				if slice[i] == key {
					sort += float64(time.Since(t0))
				}
			}

			t0 = time.Now()
			if unsorted.Search_unsorted(slice, key) {
				unsort += float64(time.Since(t0))
			}

		}
	}

	fmt.Printf("\n\nBinary average: %f\nSorted average: %f\n Unsorted: %f", bin/float64(loopsize), sort/float64(loopsize), unsort/float64(loopsize))

}

/*
func main() {
	sliceSize := 32000000
	slice := sorted.Sorted(sliceSize)
	key := slice[rand.Intn(sliceSize)]
	for i := 0; i < 32000000; i++ {

		for i := 0; i < sliceSize; i++ {

			fmt.Printf("Binary:")
			t0 := time.Now()
			if sorted.Binary_search(slice, key) {
				fmt.Printf("%f \n", float64(time.Since(t0)))
			}

			fmt.Printf("Sorted: ")
			t0 = time.Now()
			for i := 0; i < len(slice); i++ {
				if slice[i] == key {
					fmt.Printf("time: %f\n", float64(time.Since(t0)))
				}
			}

			fmt.Printf("Un-sorted: ")
			t0 = time.Now()
			if unsorted.Search_unsorted(slice, key) {
				fmt.Printf("time: %f", float64(time.Since(t0)))
			}
		}
	}

}
*/
