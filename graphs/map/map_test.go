package maps

import "testing"

func TestLookUp(t *testing.T) {
	m := NewMapsTrain()
	c := m.LookUp("Stockholm")
	for _, v := range c.Neighbors {
		println(v.City.Name)
	}

}

func TestLookUpEmpty(t *testing.T) {
	m := NewMapsTrain()
	name := "öö"
	m.LookUp(name)

	if m.LookUp(name).Name != name {
		t.Errorf("Look up did not add new city")
	}
}
