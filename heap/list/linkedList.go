package list

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type NodeR1 struct {
	tail *NodeR1
	prio int
}

type LinkedListR1 struct {
	first *NodeR1
}

func (l *LinkedListR1) Remove() (retNod *NodeR1) {
	retNod = l.first
	l.first = l.first.tail
	return
}
func MakeNodeR1(prior int) NodeR1 {
	return NodeR1{prio: prior}
}

func (l *LinkedListR1) Add(n *NodeR1) {
	if l.first == nil {
		l.first = n
		return
	}

	nIt := l.first

	if l.first.tail == nil {
		if l.first.prio < n.prio {
			l.first.tail = n
		} else {
			n.tail = l.first
			l.first = n
		}
	}

	for nIt.tail.prio <= n.prio {
		nIt = nIt.tail
	}

	n.tail = nIt.tail
	nIt.tail = n
}

type NodeA1 struct {
	tail *NodeA1
	prio int
}

type LinkedListA1 struct {
	first *NodeA1
}

func MakeNodeA1(prior int) NodeA1 {
	return NodeA1{prio: prior}
}

func (l *LinkedListA1) Add(n *NodeA1) {
	n.tail = l.first
	l.first = n
}

func (l *LinkedListA1) Remove() (retNod *NodeA1) {
	nIt := l.first
	var front *NodeA1
	retNod = l.first

	for ; nIt.tail != nil; nIt = nIt.tail {
		if nIt.tail.prio < retNod.prio {
			front = nIt
			retNod = nIt.tail
		}
	}

	front.tail = retNod.tail
	return
}

func BenchListF() {
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

		valSlice := make([]int, listSize)
		for i, _ := range valSlice {
			valSlice[i] = rand.Intn(listSize * 2)
		}

		for i := 0; i < loops; i++ {
			//append nodes to linked list through appendation
			t0 = time.Now()

			tDelta += float64(time.Since(t0))
		}

		tDelta /= float64(loops)
		fmt.Printf("%d %f\n", listSize, tDelta/1000)

	}

}
