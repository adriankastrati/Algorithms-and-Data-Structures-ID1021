package heapBinaryTree

import "testing"

func TestRemoveRoot(t *testing.T) {

	n3 := Node{Priority: 10, children: 0}
	n2 := Node{right: &n3, Priority: 9, children: 1}
	n1 := Node{Priority: 7, children: 0}
	rootPrio := 6
	r := Node{left: &n1, right: &n2, Priority: rootPrio, children: 3}

	bT := BinaryTree{root: &r}

	removed := bT.Remove()

	if removed.Priority != 6 {
		t.Errorf("prio removed: Got %d, expected %d", removed.Priority, rootPrio)
	}
	if bT.root.Priority != 7 {
		t.Errorf("not correct root removed: Got %d, expected %d", bT.root.Priority, n1.Priority)

	}
}

func TestTraverseDownNil(t *testing.T) {
	n3 := Node{Priority: 10, children: 0}
	n2 := Node{right: &n3, Priority: 9, children: 1}
	n1 := Node{Priority: 5, children: 0}

	r := Node{left: &n1, right: &n2, Priority: 6, children: 3}

	bT := BinaryTree{root: &r}
	n := bT.root

	for n != nil {
		n = n.traverseDownwards()
	}

	if n != nil {
		t.Errorf("wrong traversal, expected: nil, got %p", n)
	}
}

func TestTraverseDown(t *testing.T) {
	n3 := Node{Priority: 10, children: 0}
	n2 := Node{right: &n3, Priority: 9, children: 1}
	n1 := Node{Priority: 5, children: 0}

	r := Node{left: &n1, right: &n2, Priority: 6, children: 3}

	bT := BinaryTree{root: &r}
	n := bT.root

	node := n.traverseDownwards()

	if node != &n1 {
		t.Errorf("wrong traversel, expected: %p", &n1)
	}
}

func TestAddRoot(t *testing.T) {
	n3 := Node{Priority: 10, children: 0}
	n2 := Node{right: &n3, Priority: 9, children: 1}
	n1 := Node{Priority: 7, children: 0}

	r := Node{left: &n1, right: &n2, Priority: 6, children: 3}
	bT := BinaryTree{root: &r}
	n := Node{Priority: 1}

	bT.Add(&n)

	if bT.root.Priority != 1 {
		t.Errorf("wrong added, expected: prio %d, got %d", 1, bT.root.Priority)
	}
}

func TestAddBelowRoot(t *testing.T) {
	n3 := Node{Priority: 10, children: 0}
	n2 := Node{right: &n3, Priority: 9, children: 1}
	n1 := Node{Priority: 7, children: 0}

	r := Node{left: &n1, right: &n2, Priority: 6, children: 3}
	bT := BinaryTree{root: &r}

	n := Node{Priority: 11}

	bT.Add(&n)

	if n1.left.Priority != 11 {
		t.Errorf("wrong added, expected: %p with prio %d", &n, n.Priority)
	}

}

func TestSwap(t *testing.T) {
	n2 := &Node{Priority: 2, children: 2}
	n1 := &Node{Priority: 1, children: 1}

	n1._swap(n2)

	if n2.children != 1 || n2.Priority != 1 {
		t.Errorf("swap failed: expected children %d was %d, Priority %d was %d", n2.children, n1.children, n2.Priority, n1.Priority)
	}
	if n1.children != 2 || n1.Priority != 2 {
		t.Errorf("swap failed: expected children %d was %d, Priority %d was %d", n1.children, n2.children, n1.Priority, n2.Priority)
	}
}
