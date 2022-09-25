package list

type node struct {
	Tail *node
	Val  int
}

type LinkedList struct {
	headNode *node
}

func (l *LinkedList) GetHeadNode() *node {
	return l.headNode
}

func MakeNode(val int) node {
	n := node{}
	n.Val = val
	return n
}

func MakeLinkedList() LinkedList {
	list := LinkedList{}
	return list
}

func (n node) Next() *node {
	return n.Tail
}

func (l *LinkedList) AppendNode(val int) {
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

func (l *LinkedList) AppendList(toLink LinkedList) {

	iteratorNode := l.headNode

	for iteratorNode.Tail != nil {
		iteratorNode = iteratorNode.Tail
	}

	iteratorNode.Tail = toLink.headNode
}
