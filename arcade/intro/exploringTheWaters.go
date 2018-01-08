package intro

func alternatingSums(a []int) []int {
	first := 0
	second := 0
	for i, weight := range a {
		if i % 2 == 0 {
			first += weight
		} else {
			second += weight
		}
	}
	return []int{first, second}
}

func addBorder(picture []string) []string {
	starSize := len(picture[0]) + 2
	s := ""
	for i := 0; i < starSize; i++ {
		s += "*"
	}
	res := []string{s}
	for _, el := range picture {
		res = append(res, "*" + el + "*")
	}
	res = append(res, s)
	return res

}

func areSimilar(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	notMatchIndexes := []int{}
	for i, el := range a {

		if el != b[i] {
			notMatchIndexes = append(notMatchIndexes, i)
		}

	}
	if len(notMatchIndexes) == 0 {
		return true
	}
	if len(notMatchIndexes) != 2 {

		return false
	}

	return a[notMatchIndexes[0]] == b[notMatchIndexes[1]] && a[notMatchIndexes[1]] == b[notMatchIndexes[0]]
}

func arrayChange(inputArray []int) int {
	count := 0
	for i := 0; i < len(inputArray) - 1; i++ {
		if inputArray[i + 1] <= inputArray[i] {
			newVal := inputArray[i] + 1
			count += newVal - inputArray[i + 1]
			inputArray[i + 1] = newVal
		}
	}
	return count
}

func sortStr(s string) string {
	str := []byte(s)
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] > str[j] {
				tmp := str[i]
				str[i] = str[j]
				str[j] = tmp
			}
		}
	}
	return string(str)
}

func palindromeRearranging(inputString string) bool {
	sortedStr := sortStr(inputString)
	noPairs := 0

	for i := 1; i < len(sortedStr); i++ {
		if sortedStr[i - 1] != sortedStr[i] {
			noPairs += 1
		} else {
			i += 1
		}
	}
	if len(sortedStr) % 2 == 1 && noPairs == 0 {
		noPairs += 1
	}
	return len(sortedStr) % 2 == noPairs

}