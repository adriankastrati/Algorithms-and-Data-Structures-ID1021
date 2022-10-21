package task

import (
	"algo/heap/list"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func BenchListA1() {
	var (
		tDelta   float64
		t0       time.Time
		listSize int
		maxSize  int = 1000000
		loops    int = 100
	)
	//list A appends to list B, list A is fixed while list B increases in size

	//increase length of linked list
	for mult := 0; listSize < maxSize; mult++ {
		//restart time delta every iteration of loop

		tDelta = 0

		//size of list
		listSize = int(math.Pow(2, float64(mult)))
		listA1 := list.LinkedListA1{}

		valSlice := make([]int, listSize)
		for i := 0; i < loops; i++ {

			for i := range valSlice {
				valSlice[i] = rand.Intn(listSize * 2)
			}

			//append nodes to linked list through appendation
			t0 = time.Now()
			for _, v := range valSlice {
				listA1.Add(v)
			}

			tDelta += float64(time.Since(t0))
		}

		tDelta /= float64(loops)
		fmt.Printf("%d %f   |   ", listSize, tDelta/1000)

		tDelta = 0
		t0 = time.Now()
		for j := 0; j < listSize; j++ {
			listA1.Remove()
		}

		tDelta += float64(time.Since(t0))
		fmt.Printf("%d %f", listSize, tDelta/1000)

	}

}

func BenchListR1() {
	var (
		tDelta   float64
		t0       time.Time
		listSize int
		maxSize  int = 1000000
		loop     int = 100
	)
	//list A appends to list B, list A is fixed while list B increases in size

	//increase length of linked list
	for mult := 0; listSize < maxSize; mult++ {
		//restart time delta every iteration of loop

		tDelta = 0

		//size of list
		listSize = int(math.Pow(2, float64(mult)))
		listR1 := list.LinkedListR1{}

		valSlice := make([]int, listSize)
		// for i := 0; i < loops; i++ {
		listR1 = list.LinkedListR1{}

		for i := range valSlice {
			valSlice[i] = rand.Intn(listSize * 2)
		}

		//append nodes to linked list through appendation
		t0 = time.Now()
		for j := 0; j < listSize; j++ {
			listR1.Add(valSlice[j])
		}

		tDelta += float64(time.Since(t0))
		// }

		tDelta = float64(loop)

		fmt.Printf("%d %f", listSize, tDelta/1000)
	}

}
