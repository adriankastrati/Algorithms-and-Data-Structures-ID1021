package sort

func SelectionSort(ptrSort *[]int) {
	var (
		currMinIndex int
	)
	toSort := *ptrSort

	for i := 0; i < len(toSort)-1; i++ {
		currMinIndex = i

		for compIndex := currMinIndex; compIndex < len(toSort); compIndex++ {
			if toSort[currMinIndex] > toSort[compIndex] {
				currMinIndex = compIndex
			}
		}
		toSort[i], toSort[currMinIndex] = toSort[currMinIndex], toSort[i]
	}
}
