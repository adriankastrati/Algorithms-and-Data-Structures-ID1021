package queue

import "testing"

func TestEnQueueEmptyQueueTwoSuccess(t *testing.T) {
	QueueTwo := QueueTwo{}
	n := Node{}
	QueueTwo.EnQueue(&n)

	if QueueTwo.first != &n || QueueTwo.last != &n {
		t.Errorf("Did not add Node as last and first Node in queue")
	}
	if QueueTwo.first == nil && QueueTwo.last == nil {
		t.Errorf("Queue still empty")
	}
}

func TestEnQueueQueueTwoSuccess(t *testing.T) {
	QueueTwo := QueueTwo{}

	n1 := Node{}
	n2 := Node{}

	QueueTwo.first = &n1
	QueueTwo.last = &n1

	QueueTwo.EnQueue(&n2)

	if QueueTwo.last != &n2 {
		t.Errorf("Did not add Node as last")
	}
}

func TestRemoveLastQueueTwo(t *testing.T) {
	QueueTwo := QueueTwo{}
	n := Node{}

	QueueTwo.first = &n
	QueueTwo.last = &n

	QueueTwo.DeQueue()

	if QueueTwo.first != nil {
		t.Errorf("First Node is not nil")
	}
	if QueueTwo.last != nil {
		t.Errorf("Last Node is not nil")
	}

}

func TestDeQueueQueueTwo(t *testing.T) {
	QueueTwo := QueueTwo{}
	n3 := Node{}
	n2 := Node{}
	n2.tail = &n3
	n1 := Node{}
	n1.tail = &n2

	QueueTwo.first = &n1
	QueueTwo.last = &n3

	removed := QueueTwo.DeQueue()

	if n1.tail != nil && removed.tail != nil {
		t.Errorf("Removed Node has tail")
	}
	if QueueTwo.first != &n2 {
		t.Errorf("First Node is not second Node")
	}
	if QueueTwo.last != &n3 {
		t.Errorf("Last Node is not last Node")
	}

}

func TestAddQueueOneEmpty(t *testing.T) {
	queue := queueOne{}
	n := Node{}

	queue.EnQueue(&n)

	if queue.first != &n {
		t.Errorf("First not new Node")
	}
}

func TestAddQueueOneNotEmpty(t *testing.T) {
	queue := queueOne{}
	n1 := Node{}
	n2 := Node{}

	queue.first = &n1

	queue.EnQueue(&n2)

	nIt := queue.first

	for nIt.tail != nil {
		nIt = nIt.tail
	}

	if nIt != &n2 {
		t.Errorf("last Node not new Node")
	}
}

func TestRemoveQueueOne(t *testing.T) {
	queue := queueOne{}
	n2 := Node{}
	n1 := Node{}
	n1.tail = &n2
	queue.first = &n1

	rmN := queue.DeQueue()

	nIt := queue.first

	for nIt.tail != nil {
		nIt = nIt.tail
	}

	if nIt == rmN {
		t.Errorf("last Node still last Node")
	}
}
