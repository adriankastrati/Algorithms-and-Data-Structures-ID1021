package task

import "algo/binary-tree/binaryTree"

func TaskBinaryTreeAdd() {
	bin := binaryTree.MakeBinaryTree()
	bin.Add(2, 10)
	bin.Add(4, 20)
	bin.Add(1, 5)

	print(bin.GetRoot().GetRight().GetValue())

}
