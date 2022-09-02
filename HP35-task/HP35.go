package main

import "fmt"

type itemType int

const (
	ADD itemType = iota
	SUB
	DIV
	MUL
	VALUE
)

type item struct {
	itemType  itemType
	itemValue int
}

type calculator struct {
	expressionSlice   []item
	expressionPointer int
	stack             stack
}

func makeCalculator(exprSlice []item, stack *stack) calculator {
	calc := calculator{}
	calc.expressionPointer = 0
	calc.expressionSlice = exprSlice
	calc.stack = *stack
	return calc
}

func makeItem(itype itemType, val int) item {
	return item{itype, val}
}

func (calc *calculator) step() {

	item := calc.expressionSlice[calc.expressionPointer]

	switch item.itemType {

	case ADD:
		x := calc.stack.pop()
		y := calc.stack.pop()
		calc.stack.push(x + y)

	case SUB:
		x := calc.stack.pop()
		y := calc.stack.pop()
		calc.stack.push(x - y)

	case MUL:
		x := calc.stack.pop()
		y := calc.stack.pop()
		calc.stack.push(x * y)

	case DIV:
		x := calc.stack.pop()
		y := calc.stack.pop()
		calc.stack.push(y / x)

	case VALUE:
		calc.stack.push(item.itemValue)
	}

	calc.expressionPointer++

}

type stack struct {
	stack        []int
	stackPointer int
	stackSize    int
	stackDyn     bool
}

func makeStack(dynamicOn bool, stackSize int) (stack stack) {
	stack.stackSize = stackSize
	stack.stackDyn = dynamicOn
	stack.stackPointer = 0
	stack.stack = make([]int, stackSize)
	return
}

func (stack *stack) pop() (popvalue int) {
	if stack.stackPointer == 0 {
		panic("popped empty stack")
	}

	if stack.stackDyn && (stack.stackPointer <= len(stack.stack)/3) {
		stack.stack = stack.copyMinimizeStack()
		stack.stackSize = len(stack.stack)
	}

	stack.stackPointer--
	popvalue = stack.stack[stack.stackPointer]
	return

}

func (stack *stack) copyMinimizeStack() []int {
	oldStack := stack.stack
	newStackSize := len(oldStack) / 3

	stack.stack = make([]int, newStackSize)
	copy(stack.stack, oldStack)

	return stack.stack
}

func (stack *stack) copyExpandStack() {
	oldStack := stack.stack
	stack.stack = make([]int, stack.stackSize*2)

	for i := 0; i < stack.stackPointer; i++ {
		stack.stack[i] = oldStack[i]
	}
}

func (stack *stack) push(value int) {

	if stack.stackPointer >= stack.stackSize {
		if stack.stackDyn {
			stack.copyExpandStack()
			stack.stackSize = len(stack.stack)
			fmt.Printf("%d\n", stack.stackSize)

		} else {
			panic("pushed full stack")
		}
	}

	stack.stack[stack.stackPointer] = value
	stack.stackPointer++

}

func (c *calculator) run() int {
	for c.expressionPointer < len(c.expressionSlice) {
		c.step()
	}
	return c.stack.pop()
}

func main() {
	item1 := makeItem(VALUE, 10)
	item2 := makeItem(VALUE, 5)
	item3 := makeItem(DIV, 0)
	item4 := makeItem(VALUE, 4)
	item5 := makeItem(ADD, 5)
	item6 := makeItem(VALUE, 10)
	item7 := makeItem(VALUE, 10)
	item8 := makeItem(VALUE, 10)
	item9 := makeItem(VALUE, 10)
	item10 := makeItem(ADD, 10)
	item11 := makeItem(ADD, 10)
	item12 := makeItem(ADD, 10)
	item13 := makeItem(ADD, 10)

	itemSlice := []item{item1, item2, item3, item4, item5, item6, item7, item8, item9, item10, item11, item12, item13}

	stack1 := makeStack(true, 4)

	calc1 := makeCalculator(itemSlice, &stack1)

	fmt.Printf("%d \n", calc1.run())
	fmt.Printf("%d", calc1.stack.stackSize)
}
