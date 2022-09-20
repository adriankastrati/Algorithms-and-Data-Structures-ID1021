package task

import (
	"fmt"
	"linked-list/list"
	"math"
	"time"
)

func Task1() {
	var (
		listASize int
	)

	listA := list.MakeLinkedList()

	listB := list.MakeLinkedList()

	for mult := 5; listASize < 32000000; mult++ {
		listB = list.MakeLinkedList()

		for i := 0; i < 1000; i++ {
			listB.AppendNode(i)
		}

		listASize = int(math.Pow(2, float64(mult)))
		listA = list.MakeLinkedList()

		for i := 0; i < listASize; i++ {
			listA.AppendNode(i)
		}

		t0 := time.Now()

		listB.AppendList(&listA)

		tDelta := float64(time.Since(t0))

		fmt.Printf("%d %f\n", listASize, tDelta)
	}
}

func Task2() {
	var (
		listASize int
	)

	listA := list.MakeLinkedList()

	listB := list.MakeLinkedList()

	for mult := 5; listASize < 32000000; mult++ {
		listB = list.MakeLinkedList()

		for i := 0; i < 1000; i++ {
			listB.AppendNode(i)
		}

		listASize = int(math.Pow(2, float64(mult)))
		listA = list.MakeLinkedList()

		for i := 0; i < listASize; i++ {
			listA.AppendNode(i)
		}

		t0 := time.Now()

		listA.AppendList(&listB)

		tDelta := float64(time.Since(t0))

		fmt.Printf("%d %f\n", listASize, tDelta)
	}
}

func task3() {
	list
}
