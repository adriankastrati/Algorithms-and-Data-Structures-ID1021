package sort

func BubbleSort(ptrSort *[]int) {
	toSort := *ptrSort

	for i := 0; i < len(toSort); i++ {
		// for each element from i towards 1, swap the item found with the
		//item before it if it is smaller

		for j := i; j > 0 && j < len(toSort); j-- {

			if toSort[j-1] > toSort[j] {
				toSort[j-1], toSort[j] = toSort[j], toSort[j-1]
				continue
			}
			break
		}

	}
}
