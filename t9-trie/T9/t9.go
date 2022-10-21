package t9

import (
	"algo/t9-trie/util"
	"bufio"
	"log"
	"os"
)

func makeNode() *Node {
	nxt := make([]*Node, 27)
	return &Node{Next: nxt, Word: false}
}

type Node struct {
	Next []*Node
	Word bool
}

type T9 struct {
	Root *Node
}

func MakeT9() T9 {
	r := makeNode()
	return T9{Root: r}
}

func (t *T9) Add(word string) {
	wIt := t.Root

	for _, v := range word {
		code := util.GetCodeByChar(v)

		if wIt.Next[code] == nil {
			wIt.Next[code] = makeNode()
		}

		wIt = wIt.Next[code]
	}
	wIt.Word = true
}

func (t *T9) Search(key string) []string {

	return t.Root.search(key, "")

}

func (n *Node) search(key string, path string) []string {
	possWord := make([]string, 0)

	if len(key) > 0 {
		index := util.GetIndexByKey(key[0])
		added := false

		if n.Next[index] != nil {
			path += util.GetCharByCode(index)
			possWord = append(possWord, n.Next[index].search(key[1:], path)...)
			added = true

		}

		if n.Next[index+1] != nil {
			if added {
				path = path[:len(path)-1]
			}
			added = true

			path += util.GetCharByCode(index)
			possWord = append(possWord, n.Next[index].search(key[1:], path)...)

		}
		if n.Next[index+2] != nil {
			if added {
				path = path[:len(path)-1]
			}

			path += util.GetCharByCode(index)
			possWord = append(possWord, n.Next[index].search(key[1:], path)...)

		}
	} else {
		if n.Word {
			possWord = append(possWord, path)
			return possWord
		} else {
			return nil
		}
	}

	return possWord
}

func MakeFullT9(fileName string) T9 {
	t9 := MakeT9()

	// open file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		t9.Add(scanner.Text())

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return t9
}
