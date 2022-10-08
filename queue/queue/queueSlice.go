package queue

type QueueSlice struct {
	first, last int
	queue       []*Node
}

func MakeQueueSlice() QueueSlice {
	qS := QueueSlice{first: 0, last: 0, queue: make([]*Node, 20)}
	return qS
}

func (q *QueueSlice) IsOutBounds(index int) bool {
	if index < len(q.queue) {
		return true
	} else {
		return false
	}
}

func (q *QueueSlice) GetNode(index int) *Node {
	return q.queue[index]
}

func (q *QueueSlice) EnQueue(n *Node) {
	if q.last == q.first && q.queue[q.last] != nil {
		println("Queue is full")
		return
	} else {
		q.queue[q.last] = n
		q.last = (q.last + 1) % len(q.queue)
	}
}

func (q *QueueSlice) GetSlice() []*Node {
	return q.queue
}

func (q *QueueSlice) DeQueue() (retNod *Node) {
	if q.queue[q.first] == nil {
		println("Queue is empty")
		return nil
	}

	retNod = q.queue[q.first]
	q.queue[q.first] = nil

	q.first = (q.first + 1) % len(q.queue)
	return
}
func (q *QueueSlice) GetLast() int {
	return q.last
}

func (q *QueueSlice) GetFirst() int {
	return q.first
}

type DynQueueSlice struct {
	first, last int
	queue       []*Node
}

func (q *DynQueueSlice) GetNode(index int) *Node {
	return q.queue[index]
}

func MakeDynQueueSlice() DynQueueSlice {
	qS := DynQueueSlice{first: 0, last: 0, queue: make([]*Node, 20)}

	return qS
}
func (q *DynQueueSlice) GetQueue() []*Node {
	return q.queue
}

func (q *DynQueueSlice) EnQueue(n *Node) {

	if q.last == q.first && q.queue[q.last] != nil {
		dynS := make([]*Node, len(q.queue)*2)

		for i := 0; i < len(q.queue) && q.first != q.last; i++ {
			dynS[i] = q.queue[q.first]

			q.first = (q.first + 1) % len(q.queue)
		}

		q.first = 0
		q.last = len(q.queue)

		q.queue = dynS

	}

	q.queue[q.last] = n
	q.last = (q.last + 1) % len(q.queue)
}

func (q *DynQueueSlice) DeQueue() (retNod *Node) {
	if q.first == q.last {
		println("Queue is empty")
		return nil
	}

	retNod = q.queue[q.first]
	q.queue[q.first] = nil

	q.first = (q.first + 1) % len(q.queue)
	return
}
