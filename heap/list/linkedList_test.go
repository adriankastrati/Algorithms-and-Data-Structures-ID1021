package list

import "testing"

func TestAdd(t *testing.T) {
	listA1 := LinkedListA1{}

	listA1.Add(4)
	if listA1.First.Prio != 4 {
		t.Errorf("wrong added, expected: Prio %d, got %d", 4, listA1.First.Prio)
	}
}

func TestAddWithBeforeRoot(t *testing.T) {
	listA1 := LinkedListA1{First: &Node{Prio: 5}}

	listA1.Add(2)

	if listA1.First.Prio != 2 {
		t.Errorf("wrong added, expected: Prio %d, got %d", 2, listA1.First.Prio)
	}
}

func TestAddWithRootAfter(t *testing.T) {
	listA1 := LinkedListA1{First: &Node{Prio: 5}}

	listA1.Add(7)

	if listA1.First.Tail.Prio != 5 {
		t.Errorf("wrong added, expected: Prio %d, got %d", 7, listA1.First.Tail.Prio)
	}
}

func TestRemoveFirst(t *testing.T) {
	listA1 := LinkedListA1{First: &Node{Prio: 5}}

	listA1.Add(7)

	rN := listA1.Remove()

	if rN != 5 {
		t.Errorf("wrong added, expected: Prio %d, got %d", 5, listA1.First.Tail.Prio)
	}
}
func TestWithRemoveSecond(t *testing.T) {
	listA1 := LinkedListA1{First: &Node{Prio: 5}}

	listA1.Add(2)
	rN := listA1.Remove()

	if rN != 2 {
		t.Errorf("wrong added, expected: Prio %d, got %d", 2, rN)
	}
}
