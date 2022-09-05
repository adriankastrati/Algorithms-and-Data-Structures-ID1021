package sorted

import (
	"fmt"
	"time"
)

func Binary_search(slice []int, key int) (keyFound bool) {
	keyFound = false
	first := 0
	last := len(slice) - 1

	currIndex := (last - first) / 2
	fmt.Printf("0 Current index: %d | Current value: %d | Last: %d | First: %d | Key: %d\n", currIndex, slice[currIndex], last, first, key)

	for true {
		time.Sleep(500000000)

		if slice[currIndex] == key {
			keyFound = true
			fmt.Printf("1 Current index: %d | Current value: %d | Last: %d | First: %d | Key: %d\n", currIndex, slice[currIndex], last, first, key)
			return
		}

		if slice[currIndex] < key && currIndex < last {
			first = currIndex
			currIndex = last - ((last - first) / 2)
			fmt.Printf("2 Current index: %d | Current value: %d | Last: %d | First: %d | Key: %d\n", currIndex, slice[currIndex], last, first, key)

			continue
		}

		if slice[currIndex] > key && currIndex > first {
			last = currIndex
			currIndex = first + ((last - first) / 2)
			fmt.Printf("3 Current index: %d | Current value: %d | Last: %d | First: %d | Key: %d\n", currIndex, slice[currIndex], last, first, key)

			continue
		}

	}
	return
}

/*public static boolean binary_search(int[] array, int key) {
    int first = 0;
int last = array.length-1;

while (true) {
// jump to the middle int index = ....... ;
if (array[index] == key) { // hmm what now?
:
      }
      if (array[index] < key && index < last) {
// The index position holds something that is less than
// what we're looking for, what is the first possible page? first = ...... ;
continue;
      }
      if (array[index] > key && index > first) {
// The index position holds something that is larger than // what we're looking for, what is the last possible page? last = ...... ;
continue;
}
      // Why do we land here? What shoudl we do?
: }
    return false;
  }
*/
