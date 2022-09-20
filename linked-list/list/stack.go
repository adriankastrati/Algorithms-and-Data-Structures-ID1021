package list

type element struct {
	under *element
	val   int
}

type stack struct {
	front *element
}

func MakeElement(val int) element {
	el := element{}
	el.val = val
	return el
}

func MakeStack() stack {
	return stack{}
}

func (s *stack) Push(el element) {
	s.front = &el
}

func (s *stack) Pop() element {
	el := s.front
	s.front = s.front.under
	return *el
}
