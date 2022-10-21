package zip

import (
	"algo/hash-tables/util"
	"encoding/csv"
	"os"
	"strconv"
)

type NodeI struct {
	areaCode   int
	name       string
	population int
}

type ZipI struct {
	data []NodeI
	max  int
}

func MakeZipI(fileName string) ZipI {

	dat := make([]NodeI, 10000)

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

		n := NodeI{areaCode: code, name: line[1], population: popu}

		dat[i] = n
	}
	max := i - 1

	return ZipI{data: dat, max: max}
}

func (z *ZipI) LinearSearch(zCode int) bool {
	for i := 0; i <= z.max; i++ {
		if zCode == z.data[i].areaCode {
			return true
		}
	}
	return false
}

func (z *ZipI) BinarySearch(zCode int) (keyFound bool) {
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

func stringToInt(str string) int {
	code, e := strconv.Atoi(str)
	if e != nil {
		println("issue with ZipI %e", e)
	}
	return code
}

func removeWhiteSpace(str string) (retStr string) {
	for _, v := range str {
		if v != ' ' {
			retStr += string(v)
		}
	}
	return
}
