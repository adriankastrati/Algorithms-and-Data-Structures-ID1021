package sliceQuickSort

func partition(slice []int, start int, end int) int {
	pivEl := slice[start]

	i := start
	j := end
	for i < j {
		if slice[i] <= pivEl {
			i++
		}
		if slice[j] > pivEl {
			j--
		} else {
			slice[i], slice[j] = slice[j], slice[i]
		}
	}
	slice[j], slice[start] = slice[start], slice[j]
	return j
}

func QuickSort(slice []int, start int, end int) {
	if start < end {
		p := partition(slice, start, end)
		QuickSort(slice, start, p-1)
		QuickSort(slice, p+1, end)
	}
}
