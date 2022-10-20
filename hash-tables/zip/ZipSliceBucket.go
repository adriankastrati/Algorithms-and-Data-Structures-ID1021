package zip

import (
	"encoding/csv"
	"os"
)

type ZipSB struct {
	Data []*NodeSB
	max  int
}

type NodeSB struct {
	AreaCode   int
	Name       string
	Population int
}

func (b *ZipSB) AddNode(n NodeSB) {
	hashIndex := Hash(n.AreaCode)
	// println("hash :", hashIndex)

	for i := hashIndex; i < len(b.Data); i++ {
		// println("i:", i)
		if b.Data[i] == nil {
			b.Data[i] = &n
			b.max++
			break
		}
	}
}

func MakeZipSB(s int) ZipSB {
	nSlice := make([]*NodeSB, 2*s)

	return ZipSB{Data: nSlice, max: 0}
}

func MakeFullZipSB(fileName string) ZipSB {

	dat := make([]*NodeSB, 50000)
	zB := ZipSB{Data: dat}

	f, errOpen := os.Open(fileName)
	controlFileRead(errOpen)

	defer f.Close()

	csvReader := csv.NewReader(f)

	i := 0
	for ; ; i++ {
		line, err := csvReader.Read()

		if EOF(err) {
			break
		} else {
			controlFileRead(err)
		}

		popu := stringToInt(line[2])

		code := stringToInt(removeWhiteSpace(line[0]))

		n := NodeSB{AreaCode: code, Name: line[1], Population: popu}
		zB.AddNode(n)
	}
	zB.max = i - 1

	return zB

}

func (z *ZipSB) LookUp(zCode int) (bool, int) {
	zCodeHash := Hash(zCode)

	for i := 0; i <= z.max; i++ {
		if z.Data[i+zCodeHash] != nil {
			if z.Data[i+zCodeHash].AreaCode == zCode {
				// println("i: ", i, " hash: ", zCodeHash)
				return true, i
			}
		}
	}
	return false, 0

}

func (z *ZipSB) getZipCode() []int {
	sl := make([]int, z.max)

	for i := 0; i < z.max; i++ {
		sl[i] = z.Data[i].AreaCode

	}

	return sl
}

func (z *ZipSB) Collisions(mod int) {
	data := make([]int, mod)
	cols := make([]int, 10)

	keys := z.getZipCode()

	for i := 0; i < z.max; i++ {
		index := keys[i] % mod
		cols[data[index]]++
		data[index]++
	}

	println(mod)
	for i := 0; i < 10; i++ {
		print("\t", cols[i])
	}
	println()
}
