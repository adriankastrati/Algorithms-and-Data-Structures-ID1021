package maps

import (
	"algo/graphs/city"
)

func NaiveSearchD(fromC *city.City, toC *city.City, max int) *int {
	if max < 0 {
		return nil
	} else if fromC == toC {
		z := 0
		return &z
	}

	shortest := &max

	for _, c := range fromC.Neighbors {

		dist := NaiveSearch(c.City, toC, max-c.Distance)
		if dist != nil {

			path := *dist + c.Distance

			if path < *shortest {
				shortest = &path
			}
		}
	}

	return shortest
}

func NaiveSearch(fromC *city.City, toC *city.City, max int) *int {
	if max < 0 {
		return nil
	} else if fromC == toC {
		z := 0
		return &z
	}

	var shortest *int

	for _, c := range fromC.Neighbors {

		dist := NaiveSearch(c.City, toC, max-c.Distance)

		if dist != nil && shortest == nil {
			i := *dist + c.Distance
			shortest = &i

		} else if shortest != nil && dist != nil {
			if *shortest > (*dist + c.Distance) {
				i := *dist + c.Distance
				shortest = &i
			}
		}

		if (shortest != nil) && (max > *shortest) {
			max = *shortest
		}
	}

	return shortest
}
