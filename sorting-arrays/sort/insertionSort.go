package sort

func InsertionSort(ptrSort *[]int) {
	toSort := *ptrSort

	for i := 0; i < len(toSort); i++ {

		for j := i; j > 0 && j < len(toSort); j-- {

			if toSort[j-1] > toSort[j] {
				toSort[j-1], toSort[j] = toSort[j], toSort[j-1]
				continue
			}
			break
		}

	}
}
