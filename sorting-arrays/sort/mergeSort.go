package sort

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
