package task

import (
	doubleLinkedList "algo/double-linked-list/list"
	"algo/linked-list/linkedList"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Task1() {

	var maxSize int = 1000000
	var loop int = 100
	var op int = 500

	var (
		listSize int
		tDelta   int
		t0       time.Time
	)

	for mult := 1; listSize < maxSize; mult++ {
		tDelta = 0
		listSize = int(math.Pow(2, float64(mult)))

		for i := 0; i < loop; i++ {

			sll := doubleLinkedList.MakeLinkedList()
			var randVal int
			for k := 0; k < listSize; k++ {
				randVal += rand.Intn(10) + 1
				sll.AddV(randVal)
			}

			sllSlice := make([]*doubleLinkedList.Node, listSize)

			n := sll.GetHeadNode()
			for i := 0; i < listSize; i++ {
				sllSlice[i] = n
				n = n.Tail
			}

			for p := 0; p < op; p++ {

				rmN := sllSlice[rand.Intn(listSize)]

				t0 = time.Now()

				sll.Remove(rmN)
				sll.Add(rmN)
				tDelta += int(time.Since(t0))
			}

		}

		tDelta /= loop
		fmt.Printf("%d %f\n", listSize, float64(tDelta)/1000)
	}

}

func Tas() {

	var maxSize int = 1000000
	var loop int = 100

	var (
		listSize int
		tDelta   int
		t0       time.Time
	)

	// for mult := 9; listSize < maxSize; mult++ {
	// 	tDelta = 0
	// 	listSize = int(math.Pow(2, float64(mult)))

	// 	for i := 0; i < loop; i++ {

	// 		t0 = time.Now()

	// 		doubleLinkedList.MakeNode(10, nil, nil)

	// 		tDelta += int(time.Since(t0))

	// 	}
	// 	tDelta /= loop
	// 	fmt.Printf("%d %f\n", listSize, float64(tDelta)/1000)
	// }

	for mult := 1; listSize < maxSize; mult++ {
		tDelta = 0
		listSize = int(math.Pow(2, float64(mult)))

		slice := make([]int, listSize)
		for i, _ := range slice {
			slice[i] = rand.Intn(listSize)
		}

		for i := 0; i < loop; i++ {
			t0 = time.Now()

			for k := 0; k < listSize; k++ {
				linkedList.MakeNode(slice[k])
			}

			tDelta += int(time.Since(t0))

		}
		tDelta /= loop
		fmt.Printf("%d %f\n", listSize, float64(tDelta)/1000)
	}
	listSize = 0
	println()
	println()
	println()

	for mult := 1; listSize < maxSize; mult++ {
		tDelta = 0
		listSize = int(math.Pow(2, float64(mult)))

		slice := make([]int, listSize)
		for i, _ := range slice {
			slice[i] = rand.Intn(listSize)
		}

		for i := 0; i < loop; i++ {
			t0 = time.Now()

			for k := 0; k < listSize; k++ {
				doubleLinkedList.MakeNode(slice[k], nil, nil)
			}

			tDelta += int(time.Since(t0))

		}
		tDelta /= loop
		fmt.Printf("%d %f\n", listSize, float64(tDelta)/1000)
	}
}
