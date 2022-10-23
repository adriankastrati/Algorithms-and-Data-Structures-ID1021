package maps

import (
	"algo/graphs/city"
	"algo/hash-tables/util"
	"encoding/csv"
	"os"
)

type Maps struct {
	Cities []*city.City
	Mod    int
}

func (m *Maps) Hash(name string) int {
	hash := 7

	for _, v := range name {
		hash = (hash * 31 % m.Mod) + int(v)
	}

	return hash % m.Mod
}

func (m *Maps) LookUp(name string) *city.City {
	hashIndex := m.Hash(name)
	i := 0
	for i = hashIndex; i < len(m.Cities) && m.Cities[i] != nil; i++ {
		if m.Cities[i].Name == name {
			return m.Cities[i]
		}
	}
	m.Cities[i] = &city.City{Name: name}

	return nil
}

func (m *Maps) AddCity(c city.City) {
	hashIndex := m.Hash(c.Name)

	for i := 0; i+hashIndex < len(m.Cities); i++ {

		if m.Cities[i+hashIndex] == nil {
			m.Cities[i+hashIndex] = &c
			break
		}
		if m.Cities[i+hashIndex].Name == c.Name {
			c := city.Connection{City: c.Neighbors[0].City, Distance: c.Neighbors[0].Distance}
			m.Cities[i+hashIndex].AddConnection(c)
			break
		}
	}

}

func NewMapsTrain() Maps {
	mod := 541
	c := make([]*city.City, mod)
	maps := Maps{Mod: mod, Cities: c}

	f, errOpen := os.Open("../data/trains.csv")
	util.ControlFileRead(errOpen)

	defer f.Close()

	csvReader := csv.NewReader(f)

	for {
		line, err := csvReader.Read()

		if util.EOF(err) {
			break
		} else {
			util.ControlFileRead(err)
		}
		//use line for data
		distance := util.StringToInt(line[2])

		cityA := city.NewCity(line[1])
		connA := city.Connection{Distance: distance, City: cityA}

		cityB := city.NewCity(line[0])
		connB := city.Connection{Distance: distance, City: cityB}

		cityA.AddConnection(connB)
		cityB.AddConnection(connA)

		maps.AddCity(cityA)
		maps.AddCity(cityB)
	}
	return maps
}
