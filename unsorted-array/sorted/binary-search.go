package sorted

import (
	"time"
)

func Binary_search(slice []int, key int) (keyFound bool) {
	var (
		first int = 0
		last  int = len(slice) - 1
		mid   int = 0
	)
	keyFound = false

	if key > slice[last] {
		return
	}

	for true {
		mid = first + (last-first)/2

		if slice[mid] == key {
			keyFound = true
			return
		}

		if slice[mid] < key && mid < last {
			first = mid + 1
			continue
		}

		if slice[mid] > key && mid > first {
			last = mid - 1

			continue
		}
		break
	}

	return
}

func TimeForBinaryDuplicateSearch(keySlice []int, arrayValue []int) float64 {
	var timeTotal float64

	for keyIndex := 0; keyIndex < len(keySlice); keyIndex++ {

		time0 := time.Now()
		if Binary_search(arrayValue, keySlice[keyIndex]) {
			timeTotal += float64(time.Since(time0))
		} else {
			timeTotal += float64(time.Since(time0))

		}
	}

	return timeTotal
}

func BetterTimeForBinaryDuplicateSearch(keySlice []int, arrayValue []int) float64 {
	var (
		first   int = 0
		last    int = len(arrayValue) - 1
		mid     int = 0
		totTime float64
		time0   time.Time
	)

	for keyIndex := 0; keyIndex < len(keySlice); keyIndex++ {
		first = 0
		last = len(arrayValue) - 1
		mid = 0

		time0 = time.Now()

		if keySlice[keyIndex] > arrayValue[last] {
			totTime += float64(time.Since(time0))
			continue
		}

		for true {
			mid = first + (last-first)/2

			if arrayValue[mid] == keySlice[keyIndex] {
				totTime += float64(time.Since(time0))
			}

			if arrayValue[mid] < keySlice[keyIndex] && mid < last {
				first = mid + 1
				continue
			}

			if arrayValue[mid] > keySlice[keyIndex] && mid > first {
				last = mid - 1

				continue
			}
			break
		}
		if !(keyIndex+1 >= len(keySlice)) {
			if keySlice[keyIndex+1] > arrayValue[keyIndex+1] {
				break
			}
		}
		if !(keyIndex+1 >= len(keySlice)) {
			if keySlice[keyIndex+1] == arrayValue[keyIndex+1] {
				break
			}
		}
		if !(keyIndex+1 >= len(keySlice)) {
			if keySlice[keyIndex+1] > arrayValue[keyIndex+1] {
				break
			}
		}

	}
	return float64(time.Since(time0))

}
