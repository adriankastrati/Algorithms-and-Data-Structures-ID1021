package zip

import (
	"algo/hash-tables/util"
	"encoding/csv"
	"os"
)

type ZipB struct {
	data []Bucket
	max  int
}

type Bucket struct {
	head       *NodeB
	bucketSize int
}
type NodeB struct {
	areaCode   int
	name       string
	population int
	nextN      *NodeB
}

func (b *Bucket) addNode(n *NodeB) (i int) {
	b.bucketSize++
	if b.head == nil {
		b.head = n
		return
	}

	it := b.head
	for it.nextN != nil {
		it = it.nextN
		i++
	}
	it.nextN = n
	return
}

func MakeFullZipB(fileName string) ZipB {

	dat := make([]Bucket, 10000)

	f, errOpen := os.Open(fileName)
	util.ControlFileRead(errOpen)

	defer f.Close()

	csvReader := csv.NewReader(f)
	coll := 0
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
		hCode := util.Hash(code)

		n := NodeB{areaCode: code, name: line[1], population: popu}
		coll += (dat[hCode].addNode(&n))

	}
	println(coll)
	max := i - 1

	return ZipB{data: dat, max: max}

}

func (z *ZipB) Add(n NodeB) {
	zCodeHash := util.Hash(n.areaCode)
	z.data[zCodeHash].addNode(&n)
}

func (b *Bucket) search(code int) (bool, int) {
	firstN := b.head
	i := 0
	for firstN != nil {
		if firstN.areaCode == code {
			return true, i
		}
		firstN = firstN.nextN
		i++
	}
	return false, 0
}

func (z *ZipB) LookUp(zCode int) (bool, int) {
	bucketIndex := util.Hash(zCode)
	return z.data[bucketIndex].search(zCode)

}

func (z *ZipB) getZipCode() []int {
	sl := make([]int, z.max)

	for i := 0; i < z.max; i++ {
		sl[i] = z.data[i].head.areaCode

	}

	return sl
}

func (z *ZipB) Collisions(mod int) {
	data := make([]int, mod)
	cols := make([]int, 10)

	keys := z.getZipCode()

	for i := 0; i < z.max; i++ {
		index := keys[i] % mod
		cols[data[index]]++
		data[index]++
	}

	println(mod)
	for i := 0; i < len(cols); i++ {
		print("\t", cols[i])
	}
	println()
}
