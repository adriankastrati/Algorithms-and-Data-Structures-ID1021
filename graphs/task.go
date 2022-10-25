package task

import (
	maps "algo/graphs/map"
	"time"
)

func DFS() {
	m := maps.NewMapsTrain("graphs/data/trains.csv")

	bench("Katrineholm", "Stockholm", &m, 50)

}

func PathSearch() {
	m := maps.NewMapsTrain("graphs/data/trains.csv")

	p := maps.NewPaths()

	cityA := m.LookUp("Katrineholm")
	cityB := m.LookUp("Stockholm")

	t0 := time.Now()
	dist := p.Search(cityA, cityB, 69)
	tDelta := time.Since(t0)
	if dist != nil {
		println(cityA.Name, "-", cityB.Name)
		println(*dist, " ", tDelta)
	} else {
		println("not found or to short path")
	}
}

// func DFS() {

// 	m := maps.NewMapsTrain("graphs/data/trains.csv")

// 	bench("Stockholm", "Katrineholm", &m, 100)

// 	bench("Stockholm", "Göteborg", &m, 10000)

// 	bench("Göteborg", "Stockholm", &m, 10000)

// 	bench("Malmö", "Stockholm", &m, 10000)

// 	bench("Stockholm", "Sundvall", &m, 10000)

// 	bench("Stockholm", "Umeå", &m, 10000)

// 	bench("Göteborg", "Sundvall", &m, 10000)

// 	bench("Sundsvall", "Umeå", &m, 10000)

// 	bench("Umeå", "Göteborg", &m, 10000)

// 	bench("Göteborg", "Umeå", &m, 10000)

// }

func bench(a string, b string, m *maps.Map, max int) {

	cityA := m.LookUp(a)
	cityB := m.LookUp(b)

	t0 := time.Now()
	dist := maps.NaiveSearch(cityA, cityB, max)
	tDelta := time.Since(t0)
	if dist != nil {
		println(a, "-", b)
		println(*dist, " ", tDelta)
	} else {
		println("not found or to short path")
	}

}
