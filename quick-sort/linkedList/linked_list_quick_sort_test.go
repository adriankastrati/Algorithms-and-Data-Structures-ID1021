package linkedQuickSort

import "testing"

func TestSwap(t *testing.T) {
	list := linkedList{}

	n5 := &Node{next: nil, item: 8}
	n4 := &Node{next: n5, item: 2}
	n3 := &Node{next: n4, item: 6}
	n2 := &Node{next: n3, item: 3}
	n1 := &Node{next: n2, item: 5}

	list.first, list.last = n1, n5

	for n := list.first; n != nil; n = n.next {
		println(n.item.(int))
	}
	println("______")

}

func TestQuickSort(t *testing.T) {
	list := linkedList{}

	n4 := &Node{next: nil, item: 2}
	n3 := &Node{next: n4, item: 5}
	n2 := &Node{next: n3, item: 1}
	n1 := &Node{next: n2, item: 3}

	list.first, list.last = n1, n4

	for n := list.first; n != nil; n = n.next {
		println(n.item.(int))
	}
	println("______")

	QuickSort(list, n1, n4)
	for n := list.first; n != nil; n = n.next {
		println(n.item.(int))
	}
}

func TestRandNode(t *testing.T) {
	list := linkedList{}

	n4 := &Node{next: nil, item: 2}
	n3 := &Node{next: n4, item: 5}
	n2 := &Node{next: n3, item: 1}
	n1 := &Node{next: n2, item: 3}

	list.first, list.last = n1, n3

	list.AppendNode(n4)

	if list.last != n4 {
		t.Errorf("Did not add Node as last")
	}
}

func TestAddNode(t *testing.T) {
	list := linkedList{}

	n4 := &Node{next: nil, item: 2}
	n3 := &Node{next: n4, item: 5}
	n2 := &Node{next: n3, item: 1}
	n1 := &Node{next: n2, item: 3}

	list.first, list.last = n1, n3

	list.AppendNode(n4)

	if list.last != n4 {
		t.Errorf("Did not add Node as last")
	}
}
