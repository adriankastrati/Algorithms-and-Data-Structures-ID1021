package city

type City struct {
	Name      string
	Neighbors []Connection
}

func NewCity(city string) City {

	return City{Name: city, Neighbors: make([]Connection, 0)}
}

func (c *City) AddConnection(con Connection) {
	c.Neighbors = append(c.Neighbors, con)
}

type Connection struct {
	City     City
	Distance int
}
