package list

type LinkedListR1 struct {
	First *Node
}

func (l *LinkedListR1) Remove() (retPrio int) {
	retPrio = l.First.Prio
	l.First = l.First.Tail
	return
}
func MakeNodeR1(Prior int) Node {
	return Node{Prio: Prior}
}

func (l *LinkedListR1) Add(pr int) {
	n := &Node{Tail: nil, Prio: pr}
	if l.First == nil {
		l.First = n
		return
	}

	nIt := l.First

	for nIt != nil {
		if nIt.Prio > pr {
			n.Tail = nIt
			nIt.Front.Tail = n

		}
	}
}

type Node struct {
	Tail  *Node
	Front *Node
	Prio  int
}

type LinkedListA1 struct {
	First *Node
}

func MakeNode(Prior int) Node {
	return Node{Prio: Prior}
}

func (l *LinkedListA1) Add(pri int) {
	n := &Node{Prio: pri}
	n.Tail = l.First
	l.First = n
	n.Tail.Front = n
}

func (l *LinkedListA1) Remove() int {
	retNod := l.First
	nIt := l.First

	for nIt != nil {
		if nIt.Prio < retNod.Prio {
			retNod = nIt
		}
		nIt = nIt.Tail
	}

	if retNod.Tail == nil {
		retNod.Front.Tail = nil
	} else if retNod == l.First {
		l.First = nil
	} else if retNod.Front == nil {
		retNod.Front.Tail = retNod.Tail
	}
	return retNod.Prio
}
