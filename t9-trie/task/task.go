package task

import (
	"algo/t9-trie/t9"
)

func AddWord() {
	t9 := t9.MakeFullT9("t9-trie/data/kelly.txt")
	strings := t9.Search("324")

	for _, v := range strings {
		println(v)
	}
}
