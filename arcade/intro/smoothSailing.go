package intro

import (
	"sort"
	"strings"
)

func allLongestStrings(inputArray []string) []string {
	res := []string{}
	maxLen := 0
	for _, str := range inputArray {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}

	for _, str := range inputArray {
		if len(str) == maxLen {
			res = append(res, str)
		}
	}
	return res
}

func deleteFromStrByIndex(str string, i int) string {
	l := len(str)
	if i < 0 || i >= l {
		return str
	}
	if i == 0 {
		return str[0:]
	}
	if i == l - 1 {
		return str[:l - 2]
	}
	return str[ :i] + str[i + 1:]
}

func commonCharacterCount(s1 string, s2 string) int {
	commonChar := 0
	for _, char1 := range s1 {
		for i, char2 := range s2 {
			if char1 == char2 {
				commonChar += 1
				s2 = deleteFromStrByIndex(s2, i)
				break
			}
		}

	}
	return commonChar
}

func isLucky(n int) bool {
	numbers := []int{}
	r1, r2 := 0, 0
	for n > 0 {
		numbers = append(numbers, n % 10)
		n /= 10
	}
	halfIndex := len(numbers) / 2
	for i, el := range numbers {
		if i < halfIndex {
			r1 += el
		} else {
			r2 += el
		}
	}

	return r1 == r2

}

func sortByHeight(a []int) []int {
	nonNegative := []int{}
	index := 0
	for _, el := range a {
		if el != -1 {
			nonNegative = append(nonNegative, el)
		}
	}
	sort.Ints(nonNegative)
	for i, el := range a {
		if el != -1 {
			a[i] = nonNegative[index]
			index += 1
		}
	}
	return a

}

func reverseStr(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func reverseParentheses(s string) string {
	if !strings.ContainsAny(s, "(") {
		return s
	}
	beginIn := -1
	res := ""
	for i, el := range s {
		if el == '(' {
			beginIn = i
		}
		if el == ')' {
			res = s[:beginIn] + reverseStr(s[beginIn + 1: i]) + s[i + 1:]
			break
		}

	}
	if beginIn == -1 {
		return s
	}
	return reverseParentheses(res)
}
