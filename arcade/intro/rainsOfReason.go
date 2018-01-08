package intro

import (
	"strconv"
)

func arrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
	for i := 0; i < len(inputArray); i++ {
		if inputArray[i] == elemToReplace {
			inputArray[i] = substitutionElem
		}
	}
	return inputArray
}

func evenDigitsOnly(n int) bool {
	if n < 10 {
		return n % 2 == 0
	}
	return n % 2 == 0 && evenDigitsOnly(n / 10)
}

func variableName(name string) bool {
	if name[0] >= '0' && name[0] <= '9' {
		return false
	}
	for _, el := range name {
		if !((el >= 'a' && el <= 'z') || (el >= 'A' && el <= 'Z') || ( el >= '0' && el <= '9') || (el == '_')) {
			return false
		}
	}
	return true
}

func alphabeticShift(inputString string) string {
	res := ""
	for _, el := range inputString {
		if el == 'z' {
			res += "a"
		} else if el == 'Z' {
			res += "A"
		} else {
			res += string(el + 1)
		}
	}
	return res
}

//
func getSelIndexes(cell string) (int, int) {
	row, _ := strconv.Atoi(cell[1:])
	col := int(cell[0] - 65)
	return row - 1, col
}

func chessBoardCellColor(cell1 string, cell2 string) bool {
	board := [][]int{}

	for i := 0; i < 8; i++ {
		line := []int{}
		el := 0
		if i % 2 == 1 {
			el = 1
		}
		for j := 0; j < 8; j++ {
			line = append(line, el)
			if el == 0 {
				el = 1
			} else {
				el = 0
			}
		}
		board = append(board, line)
	}
	row1, col1 := getSelIndexes(cell1)
	row2, col2 := getSelIndexes(cell2)

	return board[row1][col1] == board[row2][col2]
}
