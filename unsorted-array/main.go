package main

import (
	"fmt"
	"math/rand"
	"unsorted-array/sorted"
)

func main() {
	sliceSize := 10
	slice := sorted.Sorted(sliceSize)
	key := slice[rand.Intn(sliceSize)]

	for i := 0; i < sliceSize; i++ {
		fmt.Printf("%d: %d\n", i, slice[i])
	}
	fmt.Printf("Starting Search for: %d\n", key)

	if sorted.Binary_search(slice, key) {
		println("found")
	} else {
		println("Not found")
	}
}
