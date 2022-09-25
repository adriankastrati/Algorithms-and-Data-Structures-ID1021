package doubleLinkedList

import "math/rand"

type LinkedList struct {
	headNode *Node
}

func MakeLinkedListRand(amount int) (l LinkedList) {

	for i := 0; i < amount; i++ {
		l.Add(rand.Intn(amount))
	}

	return
}

func (l *LinkedList) GetHeadNode() Node {
	return *l.headNode
}

func (l *LinkedList) Add(val int) {
	n := Node{}
	n.Front = nil
	n.Tail = l.headNode

	l.headNode = &n
}

func (l *LinkedList) Remove(val int) Node {
	iteratorNode := l.headNode

	for val != iteratorNode.Val {
		iteratorNode = iteratorNode.Tail
	}

	if iteratorNode.Front == nil {
		l.headNode = iteratorNode.Tail
	}
	if iteratorNode.Tail == nil {
		iteratorNode.Front.Tail = nil
	} else {
		iteratorNode.Tail.Front = iteratorNode.Front.Tail
	}
	return *iteratorNode

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
