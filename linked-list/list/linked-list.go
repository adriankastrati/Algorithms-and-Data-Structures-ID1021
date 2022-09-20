package list

type node struct {
	Tail *node
	Val  int
}

type linkedList struct {
	headNode *node
}

func (l *linkedList) GetHeadNode() *node {
	return l.headNode
}

func MakeNode(val int) node {
	n := node{}
	n.Val = val
	return n
}

func MakeLinkedList() linkedList {
	list := linkedList{}
	return list
}

func (n node) Next() *node {
	return n.Tail
}

func (l *linkedList) AppendNode(val int) {
	n := node{}
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

func (l *linkedList) AppendList(toLink linkedList) {

	iteratorNode := l.headNode

	for iteratorNode.Tail != nil {
		iteratorNode = iteratorNode.Tail
	}

	iteratorNode.Tail = toLink.headNode
}
