package binaryTree

type binaryTree struct {
	root *node
}

func (bT *binaryTree) GetRoot() *node {
	return bT.root
}

func (bT *binaryTree) Add(key int, value int) {
	if bT.root == nil {
		bT.root = &node{key, value, nil, nil}
	} else {
		n := bT.root
		newNode := MakeNode(key, value)
		n.Add(&newNode)
	}
}

func MakeBinaryTree() binaryTree {
	return binaryTree{nil}
}

type node struct {
	key         int
	val         int
	left, right *node
}

func (n *node) GetLeft() *node {
	return n.left
}

func (n *node) GetRight() *node {
	return n.right
}
func (n *node) GetValue() int {
	return n.val
}
func (n *node) Add(newNode *node) {
	if n.key < newNode.key {
		if n.right != nil {
			n.right.Add(newNode)
		} else {
			n.right = newNode
		}
	}
	if n.key > newNode.key {
		if n.left != nil {
			n.left.Add(newNode)
		} else {
			n.left = newNode
		}
	} else {
		n.val = newNode.val
	}
}

func MakeNode(key int, value int) node {
	return node{key, value, nil, nil}
}

func (bT *binaryTree) LookUp(key int) int {
	return 0
}
