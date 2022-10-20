package linkedQuickSort

import "math/rand"

type Node struct {
	next *Node
	item interface{}
}

func (n *Node) SetNext(no *Node) {
	n.next = no
}

func (n *Node) GetItem() interface{} {
	return n.item
}
func (n *Node) GetNext() *Node {
	return n.next
}

func MakeNode(ni *Node, itemi interface{}) Node {
	n := Node{}
	n.next = ni
	n.item = itemi
	return n
}

func MakeRandNode(itemi int) Node {
	n := Node{}
	n.item = rand.Intn(itemi)
	return n
}

func (l *linkedList) AppendNode(n *Node) {
	if l.first == nil {
		l.first, l.last = n, n
	} else {
		l.last.next = n
		l.last = n
	}
}

type linkedList struct {
	first *Node
	last  *Node
}

func (l *linkedList) GetFirst() *Node {
	return l.first
}
func (l *linkedList) GetLast() *Node {
	return l.last
}
func (n *linkedList) SetList(n1 *Node, n2 *Node) {
	n.first = n1
	n.last = n2
}

func MakeList() linkedList {
	return linkedList{}
}

func partition(start *Node, end *Node) (smallerN *Node) {
	pivot := start
	nIt := start.next
	smallerN = start
	for nIt != end.next {

		if pivot.item.(int) > nIt.item.(int) {
			smallerN.next.item, nIt.item = nIt.item, smallerN.next.item
			smallerN = smallerN.next

		}
		nIt = nIt.next

	}

	smallerN.item, pivot.item = pivot.item, smallerN.item
	return
}

func QuickSort(list linkedList, start *Node, end *Node) {
	if start == end || start == nil || end == nil {
		return
	}

	pivot := partition(start, end)
	QuickSort(list, start, pivot)
	QuickSort(list, pivot.next, end)
}

func (l *linkedList) Append(aL *linkedList) {
	l.last.next = aL.first
}
func (l *linkedList) PrePend(aL *linkedList) {
	aL.last.next = l.first
}
