package main

import (
	"fmt"
	"math/rand"
	"time"
)

type itemType int

const (
	ADD itemType = iota
	SUB
	DIV
	MUL
	VALUE
	MOD
	MULS
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
		if x == 0 {
			x = 1
		}
		calc.stack.push(y / x)

	case VALUE:
		calc.stack.push(item.itemValue)

	case MOD:
		x := calc.stack.pop()
		y := calc.stack.pop()
		calc.stack.push(x % y)

	case MULS:
		x := calc.stack.pop()
		y := calc.stack.pop()

		calc.stack.push(x + y - 10 + 1)
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

	if len(stack.stack) >= 10 && stack.stackDyn && (stack.stackPointer <= (len(stack.stack)/3)*2) {
		stack.copyMinimizeStack()
		stack.stackSize = len(stack.stack)
		//fmt.Printf("\nDecreased stack size: %d\n", len(stack.stack))
	}

	stack.stackPointer--
	popvalue = stack.stack[stack.stackPointer]
	return
}

/*
	func (stack *stack) copyMinimizeStack() []int {
		oldStack := stack.stack
		newStackSize := (len(oldStack) / 3) * 2

		stack.stack = make([]int, newStackSize)
		copy(stack.stack, oldStack)

		return stack.stack
	}
*/
func (stack *stack) copyMinimizeStack() {
	oldStack := stack.stack
	stack.stack = make([]int, (len(oldStack)/3)*2)

	for i := 0; i < stack.stackPointer; i++ {
		stack.stack[i] = oldStack[i]
	}
}

func (stack *stack) copyExpandStack() {
	oldStack := stack.stack
	stack.stack = make([]int, stack.stackSize*4/3)

	for i := 0; i < stack.stackPointer; i++ {
		stack.stack[i] = oldStack[i]
	}
}

func (stack *stack) push(value int) {

	if stack.stackPointer >= stack.stackSize {
		if stack.stackDyn {
			stack.copyExpandStack()
			stack.stackSize = len(stack.stack)
			//fmt.Printf("\nIncreased stack size: %d\n", len(stack.stack))

		} else {
			panic("pushed full stack")
		}
	}

	stack.stack[stack.stackPointer] = value
	stack.stackPointer++

}

func (c *calculator) run() int {
	time0 := time.Now()

	for c.expressionPointer < len(c.expressionSlice) {
		c.step()
	}

	fmt.Printf("%f", float64(time.Since(time0)))

	return c.stack.pop()
}

func generateItemOperation(amountOperation int) []item {
	itemList := make([]item, amountOperation*2-1)
	var op itemType = 0
	//	rand.Seed(time.Now().UnixNano())

	for i := 0; i < amountOperation; i++ {
		itemList[i] = makeItem(VALUE, rand.Intn(1000000)+1)
	}
	//	rand.Seed(time.Now().UnixNano())

	for i := amountOperation; i < amountOperation*2-1; i++ {
		op = itemType(rand.Intn(5))

		if op == 5 {
			op = 0
		}
		itemList[i] = makeItem(op, 0)
		op++
	}

	return itemList
}

func main() {
	/*item1 := makeItem(VALUE, 10)
	item2 := makeItem(VALUE, 5)
	item3 := makeItem(VALUE, 2)
	item4 := makeItem(VALUE, 4)
	item5 := makeItem(VALUE, 5)
	item6 := makeItem(VALUE, 10)
	item7 := makeItem(VALUE, 10)
	item8 := makeItem(VALUE, 10)
	item9 := makeItem(VALUE, 10)
	item10 := makeItem(ADD, 10)
	item11 := makeItem(ADD, 10)
	item12 := makeItem(ADD, 10)
	item13 := makeItem(ADD, 10)

	itemSlice := []item{item1, item2, item3, item4, item5, item6, item7, item8, item9, item10, item11, item12, item13}
	*/
	amount := 1000000

	itemSlicePref := generateItemOperation(amount)

	stack1 := makeStack(true, 4)
	fmt.Printf("\n\nDynamic\nTime:")
	calc1 := makeCalculator(itemSlicePref, &stack1)
	fmt.Printf("\nInput equals: %d \n", calc1.run())
	/*
		stack2 := makeStack(false, amount)
		fmt.Printf("\n\nStatic\nTime:")
		calc2 := makeCalculator(itemSlicePref, &stack2)
		fmt.Printf("\nInput equals: %d \n", calc2.run())
	*/
	/*
		sliceDigit = make([]item, 16)

		dig1 := makeItem(VALUE, 0)
		dig2 := makeItem(VALUE, 0)
		dig3 := makeItem(VALUE, 0)
		dig4 := makeItem(VALUE, 1)
		dig5 := makeItem(VALUE, 1)
		dig6 := makeItem(VALUE, 5)
		dig7 := makeItem(VALUE, 3)
		dig8 := makeItem(VALUE, 7)
		dig9 := makeItem(VALUE, 1)


		mul1 := makeItem(VALUE, 2)
		mul2 := makeItem(VALUE, 1)
		mul3 := makeItem(VALUE, 2)
		mul4 := makeItem(VALUE, 1)
		mul5 := makeItem(VALUE, 2)
		mul6 := makeItem(VALUE, 1)
		mul7 := makeItem(VALUE, 2)
		mul8 := makeItem(VALUE, 1)
		mul9 := makeItem(VALUE, 2)

		op1 := makeItem(MUL, 2)
		op2 := makeItem(MUL, 1)
		op3 := makeItem(MUL, 2)
		op4 := makeItem(MUL, 1)
		op5 := makeItem(MUL, 2)
		op6 := makeItem(MUL, 1)
		op7 := makeItem(MUL, 2)
		op8 := makeItem(MUL, 1)
		op9 := makeItem(MUL, 2)*/
}
