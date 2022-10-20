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

		valSlice := make([]*list.NodeA1, listSize)
		for i := 0; i < loops; i++ {

			for i, _ := range valSlice {
				nodeN := list.MakeNodeA1(rand.Intn(listSize * 2))
				valSlice[i] = &nodeN
			}

			//append nodes to linked list through appendation
			t0 = time.Now()
			for j := 0; j < listSize; j++ {
				listA1.Add(valSlice[j])
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
		loops    int = 100
	)
	//list A appends to list B, list A is fixed while list B increases in size

	//increase length of linked list
	for mult := 0; listSize < maxSize; mult++ {
		//restart time delta every iteration of loop

		tDelta = 0

		//size of list
		listSize = int(math.Pow(2, float64(mult)))
		listA1 := list.LinkedListR1{}

		valSlice := make([]*list.NodeR1, listSize)
		for i := 0; i < loops; i++ {

			for i, _ := range valSlice {
				nodeN := list.MakeNodeR1(rand.Intn(listSize * 2))
				valSlice[i] = &nodeN
			}

			//append nodes to linked list through appendation
			t0 = time.Now()
			for j := 0; j < listSize; j++ {
				listA1.Add(valSlice[j])
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
