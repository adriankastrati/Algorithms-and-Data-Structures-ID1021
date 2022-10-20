package heapBinaryTree

type Node struct {
	left     *Node
	right    *Node
	children int
	Priority int
}

type BinaryTree struct {
	root *Node
}

func (bT *BinaryTree) Remove() (root *Node) {
	root = bT.root

	if root.children == 1 {
		root = bT.root
		bT.root = nil
	} else if root != nil {
		bT.remove()
	}

	return

}
func (bT *BinaryTree) remove() *Node {
	return_node := bT.root
	if return_node == nil {
		return nil
	}

	var branch_to_correct *Node

	if bT.root.left == nil {
		bT.root.right = bT.root.left
		bT.root.right = nil
		return return_node
	} else if bT.root.right == nil {
		bT.root = bT.root.left
		return return_node
	}

	var new_root *Node
	if bT.root.left.Priority < bT.root.right.Priority {
		new_root = bT.root.left
		branch_to_correct = bT.root.right
	} else {
		new_root = bT.root.right
		branch_to_correct = bT.root.left
	}

	bT.root = new_root
	bT.root._add(branch_to_correct)
	return return_node
}

func (n *Node) AdoptChild(child *Node) {
	if n.left == nil {
		n.left = child
	}
	if n.right == nil {
		n.right = child
	} else {
		println("can not adopt")
	}
}

func (bT *BinaryTree) Push(increment int) int {
	bT.root.Priority += increment

	return bT.root.push()
}

func (n *Node) push() (retVal int) {
	retVal = 1
	if n.right != nil && n.left != nil {
		if n.right.Priority >= n.left.Priority && n.left.Priority < n.Priority {
			n._swap(n.left)
			retVal += n.left.push()
		} else if n.right.Priority < n.left.Priority && n.right.Priority < n.Priority {
			n._swap(n.right)
			retVal += n.right.push()
		}
	} else if n.right == nil && n.left != nil {
		if n.left.Priority < n.Priority {
			n._swap(n.left)
			retVal += n.left.push()
		}
	} else if n.right != nil && n.left == nil {
		if n.right.Priority < n.Priority {
			n._swap(n.right)
			retVal += n.right.push()
		}
	}

	return
}

func (bT *BinaryTree) Add(node *Node) int {

	if bT.root == nil {
		bT.root = node
	}
	if bT.root.Priority > node.Priority {
		bT.root._swap(node)
	}

	steps := bT.root._add(node)

	return steps
}

func (n *Node) _add(node *Node) (retVal int) {
	retVal = 1

	n.children += node.children + 1

	var moveN *Node

	if n.left == nil {
		n.left = node
		return
	}

	if n.right == nil {
		n.right = node
		return
	}

	if n.left.children <= n.right.children {
		moveN = n.left
	} else {
		moveN = n.right
	}

	if moveN.Priority > node.Priority {
		moveN._swap(node)
	}

	retVal += moveN._add(node)

	return
}

// n is removed, addN overwrites n
func (n *Node) _swap(addN *Node) {
	addN.Priority, n.Priority = n.Priority, addN.Priority
	addN.left, n.left = n.left, addN.left
	addN.right, n.right = n.right, addN.right
	addN.children, n.children = n.children, addN.children
}

func (n *Node) traverseDownwards() *Node {
	if n.left == nil && n.right == nil {
		return nil
	}
	if n.left == nil && n.right != nil {
		return n.right
	}
	if n.left != nil && n.right == nil {
		return n.left
	}

	if n.left.children < n.right.children {
		return n.left
	} else {
		return n.right
	}
}

func (bT *BinaryTree) Dive() *Node {
	n := bT.root

	for n.left != nil || n.right != nil {
		if n.right.children < n.left.children {
			n = n.right
		} else {
			n = n.left
		}
	}

	return n
}

func PromoteChild(oldNode *Node) *Node {
	var n *Node
	if oldNode.left.Priority < oldNode.right.Priority {

	} else {
		return oldNode.right
	}
	return n
}
