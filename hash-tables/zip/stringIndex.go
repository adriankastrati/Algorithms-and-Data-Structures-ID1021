package zip

import (
	"encoding/csv"
	"os"
	"strconv"
)

type NodeL struct {
	areaCode   string
	name       string
	population int
}

type ZipL struct {
	data []NodeL
	max  int
}

func MakeZipL(fileName string) ZipL {

	dat := make([]NodeL, 10000)

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

		popu, e := strconv.Atoi(line[2])
		if e != nil {
			println("issue with population %e", e)
		}

		n := NodeL{areaCode: line[0], name: line[1], population: popu}

		dat[i] = n
	}
	max := i - 1

	return ZipL{data: dat, max: max}
}

func (z *ZipL) LinearSearchL(zCode string) bool {
	for i := 0; i <= z.max; i++ {
		if zCode == z.data[i].areaCode {
			return true
		}
	}
	return false
}

func (z *ZipL) BinarySearchL(zCode string) (keyFound bool) {
	var (
		first int = 0
		last  int = len(z.data) - 1
		mid   int = 0
	)

	keyFound = false

	if zCode > z.data[mid].areaCode {
		return
	}

	for {
		mid = first + (last-first)/2

		if z.data[mid].areaCode == zCode {
			keyFound = true
			return
		}

		if z.data[mid].areaCode < zCode && mid < last {
			first = mid + 1
			continue
		}

		if z.data[mid].areaCode > zCode && mid > first {
			last = mid - 1

			continue
		}
		break
	}

	return false
}
