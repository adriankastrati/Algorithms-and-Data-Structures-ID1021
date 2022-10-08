package doubleLinkedList

import (
	"math/rand"
)

type linkedList struct {
	headNode *Node
}

func MakeLinkedList() linkedList {
	return linkedList{}
}

func MakeLinkedListRand(amount int) (l linkedList) {
	var (
		newNode Node
		rVal    int
	)

	rVal = int(rand.Intn(amount * 10))
	head := MakeNode(rVal, nil, nil)
	l.headNode = &head

	for i := 1; i < amount; i++ {
		rVal = int(rand.Intn(amount * 10))
		newNode = MakeNode(rVal, nil, nil)
		l.Add(&newNode)
	}

	return
}

func (l *linkedList) GetHeadNode() *Node {
	return l.headNode
}

func (l *linkedList) Add(n *Node) {
	n.Tail = l.headNode
	n.Front = nil

	n.Tail.Front = n
	l.headNode = n

}

func (l *linkedList) AddV(v int) {
	n := Node{}
	n.Val = v

	if l.headNode == nil {
		l.headNode = &n
		return
	}

	iteratorNode := l.headNode

	for iteratorNode.Tail != nil {
		iteratorNode = iteratorNode.Tail
	}

	iteratorNode.Tail = &n

}

func (l *linkedList) Remove(rmN *Node) {

	if rmN.Front == nil {
		l.headNode = rmN.Tail
		return
	}

	rmNF := rmN.Front
	rmNT := rmN.Tail

	if rmNF != nil && rmNT != nil {
		rmNF.Tail = rmNT
		rmNT.Front = rmNF

	} else {
		rmNF.Tail = nil
	}

}

func (l linkedList) Print() {
	n := l.headNode

	i := 0
	for n != nil {
		println("Index", i, "Value", n.Val)
		i++
		n = n.Tail
	}
	println()
}

type Node struct {
	Tail  *Node
	Front *Node
	Val   int
}

func (n *Node) GetFront() *Node {
	return n.Front
}

func (n *Node) GetTail() *Node {
	return n.Tail
}

func MakeNode(val int, tail *Node, front *Node) Node {
	n := Node{}
	n.Val = val
	n.Front = front
	n.Tail = tail

	return n
}
