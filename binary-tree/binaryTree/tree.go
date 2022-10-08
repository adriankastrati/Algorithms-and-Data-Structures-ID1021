package binaryTree

type BinaryTree struct {
	root *Node
}

func (bT *BinaryTree) GetRoot() *Node {
	return bT.root
}

func (bT *BinaryTree) LookUp(key int) retNod {
	cur := bT.root
	r := retNod{}
	r.Found = false

	if bT.root.key == key {
		r.Val = bT.root.val
		r.Found = true
		return r
	}

	for cur != nil && !r.Found {
		if cur.key == key {
			r.Val = cur.val
			r.Found = true
			return r
		}

		if cur.key < key {
			cur = cur.right
		} else {
			cur = cur.left
		}
	}

	return r
}

func (bT *BinaryTree) Add(key int, value int) {
	if bT.root != nil {
		bT.root.add(key, value)
	} else {
		bT.root = &Node{key, value, nil, nil}
	}
}

func MakeBinaryTree() BinaryTree {
	return BinaryTree{nil}
}

type Node struct {
	key         int
	val         int
	left, right *Node
}

func (n *Node) SetKey(k int) {
	n.key = k
}

type retNod struct {
	Val   int
	Found bool
}

func (n *Node) GetLeft() *Node {
	return n.left
}

func (n *Node) GetRight() *Node {
	return n.right
}
func (n *Node) GetValue() int {
	return n.val
}

func (n *Node) GetKey() int {
	return n.key
}

func (n *Node) add(key int, value int) {
	if n.key == key {
		n.val = value
	} else {
		if n.key > key {
			if n.left != nil {
				n.left.add(key, value)
			} else {
				newN := MakeNode(key, value)
				n.left = &newN
			}
		} else {
			if n.right != nil {
				n.right.add(key, value)
			} else {
				newN := MakeNode(key, value)
				n.right = &newN
			}
		}
	}
}

func MakeNode(key int, value int) Node {
	return Node{key, value, nil, nil}
}
