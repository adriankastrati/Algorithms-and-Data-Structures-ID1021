package sliceHeap

type sliceHeap struct {
	heap      []*node
	lastIndex int
}

type node struct {
	prio int
}

func makeSliceHeap() sliceHeap {
	heap := sliceHeap{lastIndex: 0}
	heap.heap = make([]*node, 0)
	return heap
}

func (sH *sliceHeap) swap(i int, j int) {
	sH.heap[i], sH.heap[j] = sH.heap[j], sH.heap[i]
}

func (sH *sliceHeap) getRoot() *node {
	return sH.heap[0]
}

func (sH *sliceHeap) Remove() (ret *node) {
	ret, sH.heap[0] = sH.heap[0], sH.heap[sH.lastIndex]
	sH.lastIndex--
	sH._sink(0)
	return
}

func (n *node) swap(toSwap *node) {
	n.prio, toSwap.prio = toSwap.prio, n.prio
}

func (sH *sliceHeap) _sink(index int) {
	if sH.heap[sH.getLeftIndex(index)] != nil && sH.heap[sH.getRightIndex(index)] != nil {
		if sH.heap[sH.getLeftIndex(index)].prio <= sH.heap[sH.getRightIndex(index)].prio {
			sH.heap[sH.getLeftIndex(index)].swap(sH.heap[index])
		} else {
			sH.heap[sH.getRightIndex(index)].swap(sH.heap[index])
		}
	} else if sH.heap[sH.getLeftIndex(index)] == nil && sH.heap[sH.getRightIndex(index)] != nil {
		sH.heap[sH.getRightIndex(index)].swap(sH.heap[index])

	} else if sH.heap[sH.getLeftIndex(index)] != nil && sH.heap[sH.getRightIndex(index)] == nil {
		sH.heap[sH.getLeftIndex(index)].swap(sH.heap[index])
	}

	return
}

func (sH *sliceHeap) getRightIndex(index int) int {
	return index*2 + 2
}

func (sH *sliceHeap) getLeftIndex(index int) int {
	return index*2 + 1
}

func (sH *sliceHeap) getParentIndex(index int) int {
	return (index - 1) / 2
}

func (sH *sliceHeap) Add(p int) {
	sH.heap[sH.lastIndex] = &node{prio: p}
	sH.lastIndex++

	sH.bubble(sH.lastIndex)
}

func (sH *sliceHeap) bubble(index int) {
	if sH.heap[sH.getParentIndex(index)].prio > sH.heap[index].prio {
		sH.heap[sH.getParentIndex(index)], sH.heap[index] = sH.heap[index], sH.heap[sH.getParentIndex(index)]
		sH.bubble(sH.getParentIndex(index))
	} else {
		return
	}

}
