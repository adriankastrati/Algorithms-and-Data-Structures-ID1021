package task

import (
	linkedQuickSort "algo/quick-sort/linkedList"
	sliceQuickSort "algo/quick-sort/slice"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Test() {
	slice := make([]int, 5)

	for i, _ := range slice {
		slice[i] = rand.Intn(35)
	}

	for i, _ := range slice {
		println(slice[i])
	}
	println("____")
	sliceQuickSort.QuickSort(slice, 2, 4)

	for k, _ := range slice {
		println(slice[k])
	}
}

func Test1() {
	list := linkedQuickSort.MakeList()

	n4 := linkedQuickSort.MakeNode(nil, 2)
	n3 := linkedQuickSort.MakeNode(&n4, 5)
	n2 := linkedQuickSort.MakeNode(&n3, 1)
	n1 := linkedQuickSort.MakeNode(&n2, 3)

	list.SetList(&n1, &n4)

	for i := list.GetFirst(); i != nil; i = i.GetNext() {
		println(i.GetItem().(int))
	}

	for i := list.GetFirst(); i != nil; i = i.GetNext() {
		println(i.GetItem().(int))
	}
}

func Bench1() {

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

		for i := 0; i < loops; i++ {
			//append nodes to linked list through appendation
			t0 = time.Now()

			listA := linkedQuickSort.MakeList()
			for k := 0; k < listSize; k++ {
				rN := linkedQuickSort.MakeRandNode(listSize * 2)
				listA.AppendNode(&rN)
			}

			tDelta += float64(time.Since(t0))
		}

		tDelta /= float64(loops)
		fmt.Printf("%d %f\n", listSize, tDelta/1000)
	}

}

func Bench2() {
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

		for i := 0; i < loops; i++ {
			//append nodes to linked list through appendation
			t0 = time.Now()

			slice := make([]int, listSize)
			for k := 0; k < listSize; k++ {
				slice[k] = rand.Intn(1000000)
			}

			tDelta += float64(time.Since(t0))
		}

		tDelta /= float64(loops)
		fmt.Printf("%d %f\n", listSize, tDelta/1000)
	}
}
