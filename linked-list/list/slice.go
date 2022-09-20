package list

func AppendSlices(sliceA []int, sliceB []int) []int {
	fullSlice := make([]int, len(sliceA)+len(sliceB))

	sliceIndex := 0
	for i := 0; i < len(sliceA); sliceIndex, i = sliceIndex+1, i+1 {
		fullSlice[sliceIndex] = sliceA[i]
	}

	for i := 0; i < len(sliceB); sliceIndex, i = sliceIndex+1, i+1 {
		fullSlice[sliceIndex] = sliceB[i]
	}

	return fullSlice
}
