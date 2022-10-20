package task

import (
	"algo/hash-tables/zip"
	"fmt"
	"time"
)

func BenchZipCode() {

	var (
		tDelta float64
		t0     time.Time
	)
	//list A appends to list B, list A is fixed while list B increases in size

	z := zip.MakeZipZ("postnr.csv")

	//restart time delta every iteration of loop
	tDelta = 0
	println("111 51")
	//append nodes to linked list through appendation
	for i := 0; i < 100; i++ {
		t0 = time.Now()

		//enter what to bench
		z.LookUp(11151)

		tDelta += float64(time.Since(t0))
	}

	fmt.Printf("Lookup:  %f\n\n", tDelta/100)
	println("984 99")
	tDelta = 0
	for i := 0; i < 100; i++ {

		t0 = time.Now()

		z.LookUp(98499)
		tDelta += float64(time.Since(t0))
	}
	fmt.Printf("Lookup:  %f\n\n", tDelta/100)

}

func BenchZipString() {

	var (
		tDelta float64
		t0     time.Time
	)
	//list A appends to list B, list A is fixed while list B increases in size

	z := zip.MakeZipL("postnr.csv")

	//restart time delta every iteration of loop
	tDelta = 0
	//size of list

	//append nodes to linked list through appendation
	t0 = time.Now()

	//enter what to bench
	println("111 51")
	z.BinarySearchL("111 51")
	tDelta = float64(time.Since(t0))

	fmt.Printf("binary:  %f\n", tDelta)

	tDelta = 0

	//append nodes to linked list through appendation
	t0 = time.Now()

	//enter what to bench
	z.LinearSearchL("111 51")
	tDelta = float64(time.Since(t0))

	fmt.Printf("linear:  %f\n\n", tDelta)
	println("984 99")

	z.BinarySearchL("984 99")
	tDelta = float64(time.Since(t0))

	fmt.Printf("binary:  %f\n", tDelta)

	tDelta = 0

	//append nodes to linked list through appendation
	t0 = time.Now()

	//enter what to bench
	z.LinearSearchL("984 99")
	tDelta = float64(time.Since(t0))

	fmt.Printf("linear:  %f\n", tDelta)
}

func BenchZipInt() {

	var (
		tDelta float64
		t0     time.Time
	)
	//list A appends to list B, list A is fixed while list B increases in size

	z := zip.MakeZipI("postnr.csv")

	//restart time delta every iteration of loop
	tDelta = 0
	//size of list

	//append nodes to linked list through appendation
	t0 = time.Now()

	//enter what to bench
	println("111 51")
	z.BinarySearch(11151)
	tDelta = float64(time.Since(t0))

	fmt.Printf("binary:  %f\n", tDelta)

	tDelta = 0

	//append nodes to linked list through appendation
	t0 = time.Now()

	//enter what to bench
	z.LinearSearch(11151)
	tDelta = float64(time.Since(t0))

	fmt.Printf("linear:  %f\n\n", tDelta)
	println("984 99")

	z.BinarySearch(98499)
	tDelta = float64(time.Since(t0))

	fmt.Printf("binary:  %f\n", tDelta)

	tDelta = 0

	//append nodes to linked list through appendation
	t0 = time.Now()

	//enter what to bench
	z.LinearSearch(98499)
	tDelta = float64(time.Since(t0))

	fmt.Printf("linear:  %f\n", tDelta)
}

func BenchCollisions() {

	z := zip.MakeZipZ("/Users/Adrian/Algorithms-and-Data-Structures-in-Go/hash-tables/data/postnr.csv")

	z.Collisions(100000000001)
}

func TestZipSB() {
	// n1 := zip.NodeSB{AreaCode: 14349, Name: "FARSTA", Population: 134}
	// n2 := zip.NodeSB{AreaCode: 22349, Name: "ESLÖV", Population: 100}

	zSB := zip.MakeFullZipSB("/Users/Adrian/Algorithms-and-Data-Structures-in-Go/hash-tables/data/postnr.csv")
	t0 := time.Now()
	f1, l1 := zSB.LookUp(10)
	t1 := time.Since(t0)
	if f1 {
		println(l1, " in ", t1, "ns")
	}

	t0 = time.Now()

	f2, l2 := zSB.LookUp(98499)
	t1 = time.Since(t0)

	if f2 {
		println(l2, " in ", t1, "ns")
	}
}

func TestZipLB() {
	lB := zip.MakeFullZipB("/Users/Adrian/Algorithms-and-Data-Structures-in-Go/hash-tables/data/postnr.csv")

	t0 := time.Now()

	f1, l1 := lB.LookUp(12152)
	t1 := time.Since(t0)

	if f1 {
		println(l1, " in ", t1, "ns")
	}

	t0 = time.Now()

	f2, l2 := lB.LookUp(98499)

	t1 = time.Since(t0)

	if f2 {
		println(l2, " in ", t1, "ns")
	}
}
