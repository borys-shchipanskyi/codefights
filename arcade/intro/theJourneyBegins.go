package intro

func add(param1 int, param2 int) int {
	return param1 + param2
}

func centuryFromYear(year int) int {
	centry := 0
	for i := 1; i <= year; i += 100 {
		centry += 1
	}
	return centry
}

func checkPalindrome(inputString string) bool {
	for i, j := 0, len(inputString) - 1; i < j; i++ {
		if inputString[i] != inputString[j] {
			return false
		}
		j -= 1
	}
	return true
}

