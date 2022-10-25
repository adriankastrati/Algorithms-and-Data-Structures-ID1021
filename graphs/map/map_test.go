package maps

import "testing"

func TestDistanceSTOSOD(t *testing.T) {
	m := NewMapsTrain("../data/trains.csv")
	c := m.LookUp("Stockholm")
	if c.Neighbors[0].Distance != 21 {
		t.Errorf("Not correct dist, expected 21 found:%d", c.Neighbors[0].Distance)
	}
}

func TestLookUp(t *testing.T) {
	m := NewMapsTrain("../data/trains.csv")
	c := m.LookUp("Stockholm")
	for _, v := range c.Neighbors {
		println(v.City.Name)
	}

}
func TestLookUpNeigbhors(t *testing.T) {
	m := NewMapsTrain("../data/trains.csv")
	c := m.LookUp("Stockholm")

	if c.Name != "Stockholm" {
		t.Errorf("Did not find Stockholm, found:%s", c.Name)
	}

}

func TestLookUpEmpty(t *testing.T) {
	m := NewMapsTrain("../data/trains.csv")
	name := "öö"
	m.LookUp(name)

	if m.LookUp(name).Name != name {
		t.Errorf("Look up did not add new city")
	}
}

func TestNaiveSearchSTOSOD(t *testing.T) {
	m := NewMapsTrain("../data/trains.csv")
	sto := m.LookUp("Stockholm")
	sod := m.LookUp("Södertälje")
	dist := NaiveSearch(sto, sod, 30)

	if *dist != 21 {
		t.Errorf("wrong distance, got %d expected %d", *dist, 21)
	}

}

func TestNaiveSearchSTOKAT(t *testing.T) {
	m := NewMapsTrain("../data/trains.csv")
	sto := m.LookUp("Stockholm")
	kat := m.LookUp("Katrineholm")
	dist := *NaiveSearch(sto, kat, 100)

	if dist != 68 {
		t.Errorf("wrong distance, got %d expected %d", dist, 68)
	}
}
