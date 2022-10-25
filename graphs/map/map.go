package maps

import (
	"algo/graphs/city"
	"algo/hash-tables/util"
	"encoding/csv"
	"os"
	"strconv"
)

type Map struct {
	Cities []*city.City
	Mod    int
}

func (m *Map) Hash(name string) int {
	hash := 7

	for _, v := range name {
		hash = (hash * 31 % m.Mod) + int(v)
	}

	return hash % m.Mod
}
func (m *Map) LookUp(name string) *city.City {
	hashCode := m.Hash(name)

	for ; hashCode < len(m.Cities); hashCode++ {

		if m.Cities[hashCode] == nil {
			c := city.NewCity(name)
			m.Cities[hashCode] = &c
			break
		}

		if m.Cities[hashCode].Name == name {
			break
		}
	}

	return m.Cities[hashCode]
}

func NewMapsTrain(filePath string) Map {
	mod := 541
	c := make([]*city.City, mod)
	Map := Map{Mod: mod, Cities: c}

	f, errOpen := os.Open(filePath)
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
		distance, _ := strconv.Atoi(line[2])

		cityA := Map.LookUp(line[0])

		cityB := Map.LookUp(line[1])

		cityA.Connect(cityB, distance)

	}
	return Map
}
