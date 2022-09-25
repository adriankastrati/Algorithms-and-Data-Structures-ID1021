package task

import (
	"double-linked-list/list"
	"math/rand"
)

func task1() {
	var (
		amountNodes       int = 100
		amountOperations  int = 50
		operationSequence []int
	)

	operationSequence = make([]int, amountOperations)

	for i := 0; i < amountOperations; i++ {
		operationSequence[i] = rand.Intn(amountNodes - 1)
	}

	doubleLinkedList := list.MakeLinkedListRand(amountNodes)

}
