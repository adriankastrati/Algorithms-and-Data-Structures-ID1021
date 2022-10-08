package iterator

import (
	"algo/binary-tree/binaryTree"
)

type treeIterator struct {
	NextN     *binaryTree.Node
	treeStack *stack
}

func (t *treeIterator) GetStackP() int {
	return t.treeStack.stackPointer
}

func MakeTreeIterator(b *binaryTree.BinaryTree) treeIterator {
	treeStack := MakeStack(true, 10)
	n := b.GetRoot()

	for n.GetLeft() != nil {
		treeStack.Push(n)
		n = n.GetLeft()
	}

	return treeIterator{n, &treeStack}
}

func (t *treeIterator) Next() (retNod *binaryTree.Node) {
	retNod = t.NextN

	if t.NextN.GetRight() != nil {
		t.NextN = t.NextN.GetRight()

		for t.NextN.GetLeft() != nil {
			t.treeStack.Push(t.NextN)
			t.NextN = t.NextN.GetLeft()
		}
		return
	}

	if t.NextN.GetRight() == nil && t.treeStack.stackPointer != 0 {
		t.NextN = t.treeStack.Pop()

		return
	}

	t.NextN = nil
	return
}

func (t *treeIterator) HasNext() bool {
	if t.NextN == nil {
		return false
	} else {
		return true
	}

}
