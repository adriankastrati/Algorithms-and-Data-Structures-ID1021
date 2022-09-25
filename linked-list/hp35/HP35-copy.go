package HP

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

func MakeCalculator(exprSlice []item, stack *stack) calculator {
	calc := calculator{}
	calc.expressionPointer = 0
	calc.expressionSlice = exprSlice
	calc.stack = *stack
	return calc
}

func MakeItem(itype itemType, val int) item {
	return item{itype, val}
}

func (calc *calculator) Step() {

	item := calc.expressionSlice[calc.expressionPointer]

	switch item.itemType {

	case ADD:
		x := calc.stack.Pop().item.itemValue
		y := calc.stack.Pop().item.itemValue
		sum := MakeElement(MakeItem(4, x+y))
		calc.stack.Push(sum)

	case SUB:
		x := calc.stack.Pop().item.itemValue
		y := calc.stack.Pop().item.itemValue
		dif := MakeElement(MakeItem(4, x-y))
		calc.stack.Push(dif)

	case MUL:
		x := calc.stack.Pop().item.itemValue
		y := calc.stack.Pop().item.itemValue
		mul := MakeElement(MakeItem(4, x*y))
		calc.stack.Push(mul)

	case DIV:
		x := calc.stack.Pop().item.itemValue
		y := calc.stack.Pop().item.itemValue
		if x == 0 {
			x = 1
		}
		if y == 0 {
			y = 1
		}

		mul := MakeElement(MakeItem(4, x/y))
		calc.stack.Push(mul)

	case VALUE:
		calc.stack.Push(MakeElement(MakeItem(4, item.itemValue)))

	case MOD:
		x := calc.stack.Pop().item.itemValue
		y := calc.stack.Pop().item.itemValue
		mod := MakeElement(MakeItem(4, x%y))

		calc.stack.Push(mod)

	case MULS:
		x := calc.stack.Pop().item.itemValue
		y := calc.stack.Pop().item.itemValue
		muls := MakeElement(MakeItem(4, x+y-10+1))

		calc.stack.Push(muls)

	}

	calc.expressionPointer++

}

func (c *calculator) Run() int {
	time0 := time.Now()

	for c.expressionPointer < len(c.expressionSlice) {
		c.Step()
	}

	fmt.Printf("%f", float64(time.Since(time0)/1000))

	return c.stack.Top.item.itemValue
}

func GenerateItemOperation(amountOperation int) []item {
	itemList := make([]item, amountOperation*2-1)
	var op itemType = 0

	for i := 0; i < amountOperation; i++ {
		itemList[i] = MakeItem(VALUE, rand.Intn(1000000)+1)
	}

	for i := amountOperation; i < amountOperation*2-1; i++ {
		op = itemType(rand.Intn(5))

		if op == 5 {
			op = 0
		}
		itemList[i] = MakeItem(op, 0)
		op++
	}

	return itemList
}
