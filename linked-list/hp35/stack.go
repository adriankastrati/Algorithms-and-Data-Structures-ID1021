package HP

type element struct {
	Under *element
	item  item
}

func (e *element) GetVal() item {
	return e.item
}

type stack struct {
	Top *element
}

func MakeElement(val item) element {
	el := element{}
	el.item = val
	return el
}

func MakeStack() stack {
	return stack{}
}

func (s *stack) Push(el element) {
	el.Under = s.Top

	s.Top = &el

}

func (s *stack) Pop() (el element) {
	el = *s.Top

	s.Top = s.Top.Under
	el.Under = nil

	return
}
