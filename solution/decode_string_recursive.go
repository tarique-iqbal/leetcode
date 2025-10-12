package solution

import (
	"unicode"
)

func decodeStringRecursive(s string) string {
	decoded, _ := dfs(s, 0)
	return decoded
}

func dfs(s string, i int) (string, int) {
	currentStr := ""
	num := 0

	for i < len(s) {
		ch := rune(s[i])

		if unicode.IsDigit(ch) {
			num = num*10 + int(ch-'0')
			i++
		} else if ch == '[' {
			subStr, nextIndex := dfs(s, i+1)
			for j := 0; j < num; j++ {
				currentStr += subStr
			}
			num = 0
			i = nextIndex
		} else if ch == ']' {
			return currentStr, i + 1
		} else {
			currentStr += string(ch)
			i++
		}
	}

	return currentStr, i
}
