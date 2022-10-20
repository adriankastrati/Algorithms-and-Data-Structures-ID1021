package t9

type Node struct {
	Next []*Node
	Word bool
}

type T9 struct {
	Root []*Node
}

func (t *T9) Add(word string) {}

func MakeTrie() T9 {
	r := make([]*Node, 27)
	return T9{Root: r}
}

func makeNode() Node {
	nxt := make([]*Node, 27)
	return Node{Next: nxt, Word: false}
}

func GetCodeByChar(word rune) int {
	switch word {
	case 'å':
		return 24
	case 'ä':
		return 25
	case 'ö':
		return 26
	}

	returnVal := int(word) - 97

	if returnVal >= 16 {
		returnVal--
	}

	if returnVal >= 22 {
		returnVal--
	}
	return returnVal
}

func GetIndexByKey(key int) int {
	return key*3 - 3
}

func GetKeyByChar(char rune) int {
	code := GetCodeByChar(char)
	return code/3 + 1

}

func (t *T9) Search(keyStroke int) []string {

	return nil
}

func GetCharByCode(code rune) string {
	switch code {
	case 24:
		return string('å')
	case 25:
		return string('ä')
	case 26:
		return string('ö')
	}

	unicodeVal := code + 97 + 2

	if unicodeVal <= 119 {
		unicodeVal--
	}
	if unicodeVal <= 113 {
		unicodeVal--
	}
	return string(unicodeVal)
}
