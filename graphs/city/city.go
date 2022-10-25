package city

type City struct {
	Name      string
	Neighbors []Connection
}

func NewCity(city string) City {

	return City{Name: city, Neighbors: make([]Connection, 0)}
}

func (cA *City) Connect(cB *City, distance int) {
	cityB_cityA := Connection{City: cB, Distance: distance}
	cA.Neighbors = append(cA.Neighbors, cityB_cityA)

	cityA_cityB := Connection{City: cA, Distance: distance}
	cB.Neighbors = append(cB.Neighbors, cityA_cityB)
}

type Connection struct {
	City     *City
	Distance int
}
