package queue

type QueueTwo struct {
	first *Node
	last  *Node
}

func (q *QueueTwo) GetFirst() *Node {
	return q.first
}

func (q *QueueTwo) GetLast() *Node {
	return q.last
}

func (q *QueueTwo) EnQueue(n *Node) {
	if q.first == nil {
		q.first = n
	} else {
		q.last.tail = n
	}
	q.last = n

}

func (q *QueueTwo) DeQueue() *Node {
	var removedN *Node

	if q.first == q.last {
		removedN = q.first
		q.first, q.last = nil, nil

	} else {
		removedN = q.first
		q.first = q.first.tail
		removedN.tail = nil

	}

	// removedN.tail = nil

	return removedN
}

type queueOne struct {
	first *Node
}

func MakeQueueOne() queueOne {
	return queueOne{}
}

func (q *queueOne) EnQueue(n *Node) {
	if q.first == nil {
		q.first = n
	}

	lastN := q.first

	for lastN.tail != nil {
		lastN = lastN.tail
	}

	lastN.tail = n
}

func (q *queueOne) DeQueue() *Node {
	removedN := q.first
	q.first = q.first.tail

	removedN.tail = nil

	return removedN
}

type Node struct {
	Item interface{}
	tail *Node
}

func (n *Node) GetItem() interface{} {
	return n.Item
}

func MakeNode(nT interface{}) Node {
	n := Node{Item: nT}

	return n
}
