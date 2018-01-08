package intro

func extractEachKth(inputArray []int, k int) []int {
	res := []int{}
	for i, el := range inputArray {
		if (i + 1) % k != 0 {
			res = append(res, el)
		}
	}
	return res
}

func firstDigit(inputString string) string {
	for _, el := range inputString {
		if el >= '0' && el <= '9' {
			return string(el)
		}
	}
	return ""
}

func differentSymbolsNaive(s string) int {
	tmp := map[string]bool{}

	for _, el := range s {
		tmp[string(el)] = true
	}
	return len(tmp)
}

func arrayMaxConsecutiveSum(inputArray []int, k int) int {
	max := 0;
	for i := 0; i <= len(inputArray) - k; i++ {
		sum := 0
		for j := i; j < i + k; j++ {
			sum += inputArray[j]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
