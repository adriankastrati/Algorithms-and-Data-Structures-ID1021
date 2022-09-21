package main

import "linked-list/task"

func main() {

	task.TimeAllocateList()

	println()
	task.TimeAllocateSlice()

	println()
	task.TimeListAOnB()

	println()
	task.TimeListBOnA()

	println()
	task.TimeSliceBOnA()
}
