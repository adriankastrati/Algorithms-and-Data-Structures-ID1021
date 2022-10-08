package task

import (
	"algo/binary-tree/binaryTree"
	"algo/binary-tree/iterator"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func MakeSliceRandom(sliceSize int) []int {
	slice := make([]int, sliceSize)

	rand.Seed(time.Now().Unix())
	for i := 0; i < sliceSize; i++ {
		slice[i] = rand.Intn(sliceSize * 100)
	}

	shuffled := rand.Perm(len(slice))
	for i := 0; i < len(shuffled); i++ {
		shuffled[i] = slice[shuffled[i]]
	}

	return shuffled
}

func TaskBinaryTreeAdd() {

	bin := binaryTree.MakeBinaryTree()

	bin.Add(10, 20)
	bin.Add(4, 20)
	bin.Add(12, 5)
	bin.Add(3, 20)
	bin.Add(1, 20)
	bin.Add(11, 5)

	it := iterator.MakeTreeIterator(&bin)

	for it.HasNext() {
		println(it.NextN.GetKey())
		it.Next()
	}

}

func TaskSearchKey() {
	bin := binaryTree.MakeBinaryTree()
	bin.Add(10, 20)
	bin.Add(4, 20)
	bin.Add(12, 5)
	it := iterator.MakeTreeIterator(&bin)

	for it.HasNext() {
		println(it.NextN.GetKey())
		it.Next()
	}

	r := bin.LookUp(1)
	if !r.Found {
		println("not found")
	} else {
		println("found", r.Val)
	}

}

func TraverseTree() {
	tree := binaryTree.MakeBinaryTree()
	tree.Add(1, 0)
	tree.Add(10, 100)
	tree.Add(2, 102)
	tree.Add(4, 107)
	tree.Add(9, 101)
	tree.Add(3, 108)

	it := iterator.MakeTreeIterator(&tree)

	for i := 0; i < 3; i++ {
		println(it.NextN.GetKey())
		it.Next()
	}

	tree.Add(15, 107)
	tree.Add(0, 101)
	tree.Add(7, 108)

	for it.HasNext() {
		println(it.NextN.GetKey())
		it.Next()
	}
}

func BenchLookUp() {
	var (
		maxSize      = 64000000
		tDelta       float64
		t0           time.Time
		amountNodes  int
		loops        = 100
		randKeySlice []int
		tMin         float64
	)
	//list A appends to list B, list A is fixed while list B increases in size

	//increase length of linked list
	for mult := 1; amountNodes < maxSize; mult++ {
		//restart time delta every iteration of loop
		tDelta = 0

		tMin = math.MaxFloat64

		//size of list
		amountNodes = 100 * mult

		for i := 0; i < loops; i++ {
			//append nodes to linked list through appendation
			randKeySlice = MakeSliceRandom(amountNodes)

			tree := binaryTree.MakeBinaryTree()

			for k := 0; k < amountNodes; k++ {
				tree.Add(randKeySlice[k], randKeySlice[k])
			}

			randKeySlice = MakeSliceRandom(amountNodes)

			t0 = time.Now()

			for k := 0; k < amountNodes; k++ {
				tree.LookUp(randKeySlice[k])
			}

			tDelta = float64(time.Since(t0))

			if tDelta < float64(tMin) {
				tMin = tDelta
			}
		}

		tDelta /= float64(loops)
		fmt.Printf("%d %f\n", amountNodes, tMin/1000)
	}

}
