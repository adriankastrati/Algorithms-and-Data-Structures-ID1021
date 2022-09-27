package linkedList

import "fmt"

type Node struct {
	Tail *Node
	Val  int
}

type LinkedList struct {
	headNode *Node
}

func (l *LinkedList) GetHeadNode() *Node {
	return l.headNode
}

func MakeNode(val int) Node {
	n := Node{}
	n.Val = val
	return n
}

func MakeLinkedList() LinkedList {
	list := LinkedList{}
	return list
}

func (l *LinkedList) AddNode(n *Node) {
	n.Tail = l.headNode
	l.headNode = n
}

func (l *LinkedList) AppendNode(val int) {
	n := Node{}
	n.Val = val

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

func (l *LinkedList) AppendList(toLink LinkedList) {
	iteratorNode := l.headNode

	for iteratorNode.Tail != nil {
		iteratorNode = iteratorNode.Tail
	}

	iteratorNode.Tail = toLink.headNode
}

func (l *LinkedList) Print() {
	n := l.headNode
	i := 0
	for n != nil {
		fmt.Printf("Index %d Value %d\n", i, n.Val)
		i++
		n = n.Tail
	}
}

func (l *LinkedList) Remove(rmN *Node) {

	if l.headNode.Val == rmN.Val {
		l.headNode = l.headNode.Tail
		return
	}

	rmNF := l.GetHeadNode()

	for rmN.Val != rmNF.Tail.Val {
		rmNF = rmNF.Tail
	}

	if rmNF.Tail.Tail == nil {
		rmNF.Tail = nil
	} else {
		rmNF.Tail = rmNF.Tail.Tail
	}

}
