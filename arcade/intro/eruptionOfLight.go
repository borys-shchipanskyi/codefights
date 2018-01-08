package intro

import (
	"sort"
	"strings"
)

func isBeautifulString(inputString string) bool {
	letters := [26]int{}
	for i := 0; i < len(inputString); i++ {
		index := int(inputString[i] - 'a')
		letters[index] += 1
	}
	for i := 0; i < 25; i++ {
		if letters[i] < letters[i + 1] {
			return false
		}
	}
	return true
}

func findEmailDomain(address string) string {
	atI := 0
	for i := len(address) - 1; i > 0; i-- {
		if address[i] == '@' {
			atI = i
			break
		}
	}

	return address[atI + 1: ]
}

//
func isPalindrome(s string) bool {
	for i, j := 0, len(s) - 1; i <= j; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

func strReverse(s string) string {
	res := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return string(res)
}

func buildPalindrome(st string) string {
	if isPalindrome(st) {
		return st
	}
	for i := 0; i < len(st); i++ {
		newStr := st + strReverse(st[:i])
		//fmt.Println(newStr)
		if isPalindrome(newStr) {
			return newStr
		}
	}
	return ""
}
//
func electionsWinners(votes []int, k int) int {
	max := votes[0]
	for i := 1; i < len(votes); i++ {
		if votes[i] > max {
			max = votes[i]
		}
	}
	possibleToWin := []int{}
	for _, el := range votes {
		if k == 0 {
			if el == max {
				possibleToWin = append(possibleToWin, el + k)
			}
		} else if el + k > max {
			possibleToWin = append(possibleToWin, el + k)
		}
	}
	sort.Ints(possibleToWin)
	if k == 0 {
		for i := 1; i < len(possibleToWin); i++ {
			if possibleToWin[i - 1] == possibleToWin[i] {
				return 0
			}
		}
	}

	return len(possibleToWin)

}

//
func isPairValid(pair string) bool {
	if len(pair) != 2 {
		return false
	}
	if (pair[0] >= '0' && pair[0] <= '9') || (pair[0] >= 'A' && pair[0] <= 'F') {
		if (pair[1] >= '0' && pair[1] <= '9') || (pair[1] >= 'A' && pair[1] <= 'F') {
			return true
		}
	}

	return false
}

func isMAC48Address(inputString string) bool {
	pairs := strings.Split(inputString, "-")
	if len(pairs) != 6 {
		return false
	}
	for _, pair := range pairs {
		if !isPairValid(pair) {
			return false
		}
	}
	return true
}
