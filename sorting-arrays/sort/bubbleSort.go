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

func MergeSort(sortSlice []int) []int {
	if len(sortSlice) == 1 {
		return sortSlice
	} else {
		return merge(
			MergeSort(sortSlice[:len(sortSlice)/2]),
			MergeSort(sortSlice[len(sortSlice)/2:]),
		)
	}

}

func merge(aSlice []int, bSlice []int) []int {
	var (
		bSliceIndex int = 0
		aSliceIndex int = 0
		mergedIndex int = 0
	)

	mergedSlice := make([]int, len(bSlice)+len(aSlice))

	for ; len(aSlice) > aSliceIndex && len(bSlice) > bSliceIndex; mergedIndex++ {

		if aSlice[aSliceIndex] > bSlice[bSliceIndex] {
			mergedSlice[mergedIndex] = bSlice[bSliceIndex]

			bSliceIndex++
			continue
		}
		if aSlice[aSliceIndex] < bSlice[bSliceIndex] {
			mergedSlice[mergedIndex] = aSlice[aSliceIndex]
			aSliceIndex++
			continue
		}

		if aSlice[aSliceIndex] == bSlice[bSliceIndex] {
			mergedSlice[mergedIndex] = aSlice[aSliceIndex]
			aSliceIndex++

			mergedIndex++

			mergedSlice[mergedIndex] = bSlice[bSliceIndex]
			bSliceIndex++
		}

	}

	for len(aSlice) > aSliceIndex {
		mergedSlice[mergedIndex] = aSlice[aSliceIndex]
		aSliceIndex++
		mergedIndex++
	}

	for len(bSlice) > bSliceIndex {
		mergedSlice[mergedIndex] = bSlice[bSliceIndex]
		bSliceIndex++
		mergedIndex++
	}

	return mergedSlice
}
