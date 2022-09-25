package linkedList

type element struct {
	Under *element
	val   int
}

func (e *element) GetVal() int {
	return e.val
}

type stack struct {
	Top *element
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
	el.Under = s.Top

	s.Top = &el

}

func (s *stack) Pop() (el element) {
	el = *s.Top

	s.Top = s.Top.Under
	el.Under = nil

	return
}
