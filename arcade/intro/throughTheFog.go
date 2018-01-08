package intro

import "fmt"

func circleOfNumbers(n int, firstNumber int) int {
	if firstNumber >= n / 2 {
		return n / 2 + firstNumber - n
	}
	return n / 2 + firstNumber
}
func depositProfit(deposit int, rate int, threshold int) int {
	counter := 0
	s := float32(deposit)
	pct := float32(rate) / float32(100)
	for s < float32(threshold) {
		s = s + s * pct
		fmt.Println(s)
		counter += 1
	}
	return counter
}
func absoluteValuesSumMinimization(a []int) int {
	l := len(a)
	if l % 2 == 1 {
		return a[l / 2]
	}
	return a[l / 2 - 1]
}

//
func getDiffCount(s1, s2 string) int {
	diffCounter := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffCounter += 1
		}
	}
	return diffCounter
}
func getPossibleNext(el string, inputArray []string) []string {
	res := []string{}
	for i := 0; i < len(inputArray); i++ {
		if getDiffCount(el, inputArray[i]) == 1 {
			res = append(res, inputArray[i])
		}
	}
	//fmt.Println("possible: ", el, res)
	return res
}

func getStrIndexInArray(el string, array []string) int {
	for i, str := range array {
		if str == el {
			return i
		}
	}
	return -1
}

func deleteFromSliceByIndex(index int, slice []string) []string {
	res := []string{}
	for i, el := range slice {
		if i != index {
			res = append(res, el)
		}
	}
	return res
}

func isPossibleToRearrangement(root string, array []string) bool {
	if len(array) == 1 {
		return getDiffCount(root, array[0]) == 1
	}
	for _, el := range getPossibleNext(root, array) {
		elIndex := getStrIndexInArray(el, array)
		newArr := deleteFromSliceByIndex(elIndex, array)
		fmt.Println(root, array, newArr, el)
		if isPossibleToRearrangement(el, newArr) {
			return true
		}
	}
	return false
}

func stringsRearrangement(inputArray []string) bool {

	for i, el := range inputArray {
		posNext := getPossibleNext(el, inputArray)
		if len(posNext) > 0 {
			if isPossibleToRearrangement(el, deleteFromSliceByIndex(i, inputArray)) {
				return true
			}

		}
	}
	return false
}
