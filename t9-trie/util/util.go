package util

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

func GetIndexByKey(key byte) int {
	var index int

	switch key {
	case '1':
		index = 0
	case '2':
		index = 3
	case '3':
		index = 6
	case '4':
		index = 9
	case '5':
		index = 12
	case '6':
		index = 15
	case '7':
		index = 18
	case '8':
		index = 21
	case '9':
		index = 24
	}

	return index
}

func GetKeyByChar(char rune) int {
	code := GetCodeByChar(char)
	return code/3 + 1

}

func GetCharByCode(code int) string {
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
	return string(rune(unicodeVal))
}
