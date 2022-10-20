package zip

import (
	"encoding/csv"
	"os"
)

type ZipZ struct {
	data []NodeI
	max  int
}

func MakeZipZ(fileName string) ZipZ {

	dat := make([]NodeI, 100000)

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

		n := NodeI{areaCode: code, name: line[1], population: popu}

		dat[code] = n
	}
	max := i - 1

	return ZipZ{data: dat, max: max}
}

func (z *ZipZ) LookUp(zCode int) bool {
	return z.data[zCode].areaCode == zCode

}

func (z *ZipZ) getZipCode() []int {
	sl := make([]int, z.max)

	for i := 0; i < z.max; i++ {
		sl[i] = z.data[i].areaCode

	}

	return sl
}

func (z *ZipZ) Collisions(mod int) {
	data := make([]int, mod)
	cols := make([]int, 100)

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
