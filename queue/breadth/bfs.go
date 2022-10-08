package breadth

import (
	"algo/binary-tree/binaryTree"
	"algo/queue/queue"
)

type treeIterator struct {
	//	NextN *binaryTree.Node
	queue *queue.QueueTwo
}

func (t *treeIterator) GetQueue() *queue.QueueTwo {
	return t.queue
}

func MakeTreeIterator(b *binaryTree.BinaryTree) treeIterator {
	tIt := treeIterator{}
	tIt.queue = &queue.QueueTwo{}

	firstQ := queue.Node{Item: b.GetRoot()}
	tIt.queue.EnQueue(&firstQ)

	return tIt
}

func (t *treeIterator) Next() (retNod *binaryTree.Node) {
	retNod = t.queue.DeQueue().GetItem().(*binaryTree.Node)

	if retNod.GetLeft() != nil {
		qNL := queue.MakeNode(retNod.GetLeft())
		t.queue.EnQueue(&qNL)
	}

	if retNod.GetRight() != nil {
		qNR := queue.MakeNode(retNod.GetRight())
		t.queue.EnQueue(&qNR)
	}

	return
}

func (t *treeIterator) HasNext() bool {
	if t.queue.GetLast() == nil && t.queue.GetFirst() == nil {
		return false
	} else {
		return true
	}
}
