package task

import (
	"algo/binary-tree/binaryTree"
	"algo/queue/queue"
)

func makeThreeLayerTree() binaryTree.BinaryTree {
	bT := binaryTree.MakeBinaryTree()
	bT.Add(10, 10)

	bT.Add(5, 5)
	bT.Add(15, 15)

	bT.Add(12, 10)
	bT.Add(17, 5)
	bT.Add(7, 7)
	bT.Add(2, 2)
	return bT
}

func SearchTree() {
	// bT := makeThreeLayerTree()

	// queIt := breadth.MakeTreeIterator(&bT)

	q := queue.MakeQueueSlice()

	println("adding nodes")
	for i := 0; i < 21; i++ {
		n := binaryTree.Node{}
		n.SetKey(i)
		println("key of node:", n.GetKey())

		qN := queue.Node{Item: &n}
		q.EnQueue(&qN)

		println("bL", q.GetLast(), "bF", q.GetFirst())
		println("aL", q.GetLast(), "bF", q.GetFirst())
		println()

	}

	println("-----")

	qN := q.DeQueue().GetItem().(*binaryTree.Node)
	println(qN.GetKey())

	n := binaryTree.Node{}
	n.SetKey(50)
	en := queue.Node{Item: &n}
	q.EnQueue(&en)

	println("Controlling nodes:...")
	for i := 0; i < 20; i++ {
		qN := q.DeQueue()
		if qN != nil {
			print(qN.GetItem().(*binaryTree.Node).GetKey())
		}

		// nf := binaryTree.Node{}
		// nf.SetKey(55)
		// pn := queue.Node{Item: &nf}
		// q.EnQueue(&pn)

		// println("Controlling nodes:...")
		// for i := 0; i < 19; i++ {
		// 	qN := q.DeQueue().GetItem().(*binaryTree.Node)
		// 	println(qN.GetKey())
		// }
	}
}
