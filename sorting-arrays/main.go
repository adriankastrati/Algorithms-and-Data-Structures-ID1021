package main

import (
	"fmt"
	"math/rand"
	"sorting-arrays/sort"
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

func main() {

	slice := MakeSlice(6)

	fmt.Println(slice)

	slice = sort.MergeSort(slice)

}
