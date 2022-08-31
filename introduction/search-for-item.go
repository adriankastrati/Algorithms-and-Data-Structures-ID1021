package main

import (
	"fmt"
	"math/rand"
	"time"
)

func search_for_item(key_length int, value_length int, amount_operations int) float64 {
	var (
		time_total  time.Duration
		slice_key       = make([]int, key_length)
		slice_value     = make([]int, value_length)
		sum         int = 0
	)

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
					break
				}
			}
		}

		time_total += time.Now().Sub(time0)
	}
	return float64(time_total) / float64(amount_operations)
}

func main() {
	for mult := 1; mult < 300; mult++ {
		var keys, values, sample_size = 11 * mult, 10 * mult, 100
		var time_per_search = search_for_item(keys, values, sample_size)

		fmt.Printf("%d %f\n", values, time_per_search)

	}
}
