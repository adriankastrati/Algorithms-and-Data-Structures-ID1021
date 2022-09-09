package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Find_duplicates(key_length int, value_length int, amount_operations int) (time_average float64, sum int) {
	var (
		time_total  time.Duration
		slice_key   = make([]int, key_length)
		slice_value = make([]int, value_length)
	)

	sum = 0

	rand.Seed(time.Now().UTC().UnixNano())

	for j := 0; j < amount_operations; j++ {
		for i := 0; i < key_length; i++ {
			slice_key[i] = rand.Intn(value_length * 10)
		}

		for i := 0; i < value_length; i++ {
			slice_value[i] = rand.Intn(value_length * 10)
		}

		time0 := time.Now()

		for key_index := 0; key_index < key_length; key_index++ {
			current_key := slice_key[key_index]

			for value_index := 0; value_index < value_length; value_index++ {
				if slice_value[value_index] == current_key {
					sum++
				}
			}
		}

		time_total += time.Now().Sub(time0)
	}

	time_average = float64(time_total) / float64(amount_operations)

	return
}

func main() {
	var (
		amount_duplicates   int
		time_for_search     float64
		i                   int = 0
		amount_arrays_timer int
		sliceSize           int
	)

	for mult := 1; mult < 300; mult++ {
		var sample_size = 100
		sliceSize = int(math.Pow(2, float64(mult)))

		time_per_search, sum := Find_duplicates(sliceSize*5, sliceSize, sample_size)

		amount_duplicates += sum

		if i == 0 {
			time_for_search += float64(time_per_search) * float64(sample_size)
			if time_for_search >= 60000000000 {
				amount_arrays_timer = values
			}
			i = 1
		}
		fmt.Printf("%d %d", amount_duplicates, amount_arrays_timer)

	}

}
