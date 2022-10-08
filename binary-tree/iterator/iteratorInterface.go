package iterator

import BinaryTree "algo/binary-tree/binaryTree"

type Iterator interface {
	HasNext() bool
	GetNext() *BinaryTree.Node
}
