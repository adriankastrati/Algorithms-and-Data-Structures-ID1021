package util

import (
	"flag"
	"io"
	"strconv"
)

func Hash(key int) int {
	return key % 10000
}

func StringToInt(str string) int {
	code, e := strconv.Atoi(str)
	if e != nil {
		println("issue with ZipI %e", e)
	}
	return code
}

func ControlFileRead(e error) {
	if e != nil {
		panic(e)
	}
}

func EOF(e error) bool {
	if e == io.EOF {
		return true
	} else {
		return false
	}
}

func GetPath(str string) *string {
	fptr := flag.String("fpath", str, "file path to read from")
	flag.Parse()
	return fptr
}
