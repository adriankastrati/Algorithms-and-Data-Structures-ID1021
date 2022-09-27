package task

import (
	"algo/linked-list/linkedList"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func task1() {

	var maxSize int = 1000000
	var loop int = 100
	var op int = 500

	var (
		listSize int
		tDelta   int
		t0       time.Time
		sll      linkedList.LinkedList
	)

	for mult := 9; listSize < maxSize; mult++ {
		tDelta = 0
		listSize = int(math.Pow(2, float64(mult)))

		for i := 0; i < loop; i++ {

			sll = linkedList.MakeLinkedList()
			var randVal int
			for k := 0; k < listSize; k++ {
				randVal += rand.Intn(10) + 1
				sll.AppendNode(randVal)
			}

			sllSlice := make([]*linkedList.Node, listSize)

			n := sll.GetHeadNode()
			for i := 0; i < listSize; i++ {
				sllSlice[i] = n
				n = n.Tail
			}

			for p := 0; p < op; p++ {

				rmN := sllSlice[rand.Intn(listSize)]

				t0 = time.Now()

				sll.Remove(rmN)
				sll.AddNode(rmN)
				tDelta += int(time.Since(t0))
			}

		}

		tDelta /= loop
		fmt.Printf("%d %f\n", listSize, float64(tDelta)/1000)
	}

}
