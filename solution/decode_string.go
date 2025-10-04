package solution

import (
	"unicode"
)

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	currentNum := 0
	currentStr := ""

	for _, ch := range s {
		if unicode.IsDigit(ch) {
			currentNum = currentNum*10 + int(ch-'0')
		} else if ch == '[' {
			numStack = append(numStack, currentNum)
			strStack = append(strStack, currentStr)
			currentNum = 0
			currentStr = ""
		} else if ch == ']' {
			n := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prevStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			temp := ""
			for i := 0; i < n; i++ {
				temp += currentStr
			}

			currentStr = prevStr + temp
		} else {
			currentStr += string(ch)
		}
	}

	return currentStr
}
