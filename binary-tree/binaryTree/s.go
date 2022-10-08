package binaryTree

type stack struct {
	stack        []*Node
	stackPointer int
	stackSize    int
	stackDyn     bool
}

func MakeStack(dynamicOn bool, stackSize int) (stack stack) {
	stack.stackSize = stackSize
	stack.stackDyn = dynamicOn
	stack.stackPointer = 0
	stack.stack = make([]*Node, stackSize)
	return
}

func (stack *stack) Pop() (popNode *Node) {
	if stack.stackPointer == 0 {
		panic("popped empty stack")
	}

	if len(stack.stack) >= 10 && stack.stackDyn && (stack.stackPointer <= (len(stack.stack)/3)*2) {
		stack.copyMinimizeStack()
		stack.stackSize = len(stack.stack)
		//fmt.Printf("\nDecreased stack size: %d\n", len(stack.stack))
	}

	stack.stackPointer--
	popNode = stack.stack[stack.stackPointer]
	return
}

func (stack *stack) copyMinimizeStack() {
	oldStack := stack.stack
	stack.stack = make([]*Node, (len(oldStack)/3)*2)

	for i := 0; i < stack.stackPointer; i++ {
		stack.stack[i] = oldStack[i]
	}
}

func (stack *stack) copyExpandStack() {
	oldStack := stack.stack
	stack.stack = make([]*Node, stack.stackSize*4/3)

	for i := 0; i < stack.stackPointer; i++ {
		stack.stack[i] = oldStack[i]
	}
}

func (stack *stack) Push(node *Node) {

	if stack.stackPointer >= stack.stackSize {
		if stack.stackDyn {
			stack.copyExpandStack()
			stack.stackSize = len(stack.stack)

		} else {
			panic("pushed full stack")
		}
	}

	stack.stack[stack.stackPointer] = node
	stack.stackPointer++

}
