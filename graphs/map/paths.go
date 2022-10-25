package maps

import "algo/graphs/city"

type Paths struct {
	path   []*city.City
	stackP int
}

func NewPaths() Paths {
	return Paths{path: make([]*city.City, 54), stackP: 0}

}

func (p *Paths) Search(fromC *city.City, toC *city.City, max int) *int {
	if max < 0 {
		return nil
	} else if fromC == toC {
		z := 0
		return &z
	}

	for _, p := range p.path {
		if p == fromC {
			return nil
		}

	}

	p.path[p.stackP] = fromC
	p.stackP++

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

	p.stackP--
	p.path[p.stackP] = nil
	return shortest
}
