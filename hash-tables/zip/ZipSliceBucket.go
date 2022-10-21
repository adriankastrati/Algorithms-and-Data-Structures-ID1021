package zip

import (
	"algo/hash-tables/util"
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

func (b *ZipSB) AddNode(n NodeSB) int {
	hashIndex := util.Hash(n.AreaCode)

	for i := 0; i+hashIndex < len(b.Data); i++ {
		if b.Data[i+hashIndex] == nil {
			b.Data[i+hashIndex] = &n
			b.max++
			return i

		}
	}
	return 0
}

func MakeZipSB(s int) ZipSB {
	nSlice := make([]*NodeSB, 2*s)

	return ZipSB{Data: nSlice, max: 0}
}

func MakeFullZipSB(fileName string) ZipSB {

	dat := make([]*NodeSB, 13000)

	zB := ZipSB{Data: dat}

	f, errOpen := os.Open(fileName)
	util.ControlFileRead(errOpen)

	defer f.Close()

	csvReader := csv.NewReader(f)

	i := 0
	for ; ; i++ {
		line, err := csvReader.Read()

		if util.EOF(err) {
			break
		} else {
			util.ControlFileRead(err)
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
	zCodeHash := util.Hash(zCode)
	var found bool

	i := 0
	for z.Data[i+zCodeHash] != nil {
		if z.Data[i+zCodeHash].AreaCode == zCode {
			found = true
			break
		}
		i++
	}

	return found, i

}

func (z *ZipSB) getZipCode() []int {
	sl := make([]int, z.max)

	for i := 0; i < z.max; i++ {
		sl[i] = z.Data[i].AreaCode

	}

	return sl
}
