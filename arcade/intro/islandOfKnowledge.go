package intro

import (
	"strings"
	"strconv"
	"sort"
	"fmt"
)

func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	return (yourLeft == friendsLeft && yourRight == friendsRight) || (yourLeft == friendsRight && yourRight == friendsLeft)
}

func arrayMaximalAdjacentDifference(inputArray []int) int {
	maxDiff := 0
	for i := 1; i < len(inputArray); i++ {
		res := inputArray[i] - inputArray[ i - 1]
		if res < 0 {
			res *= -1
		}
		if maxDiff < res {
			maxDiff = res
		}
	}
	return maxDiff
}

func isIPv4Address(inputString string) bool {
	s := strings.Split(inputString, ".")
	if len(s) != 4 {
		return false
	}
	for _, el := range s {
		if el == "" {
			return false
		}
		val, err := strconv.Atoi(el)
		if err != nil {
			return false
		}
		if val < 0 || val > 255 {
			return false
		}
	}
	return true
}

//

//func isContain(arr *[]int, val int) bool {
//	for _, el := range *arr {
//		if el == val {
//			return true
//		}
//	}
//	return false
//}

func isValid(arr*[]int, val int) bool {
	for _, el := range *arr {
		if el % val == 0 {
			return false
		}
	}
	return true
}

func avoidObstacles(inputArray []int) int {
	if len(inputArray) == 0 {
		return 0
	}
	sort.Ints(inputArray)
	last := len(inputArray) - 1
	for i := 1; i <= inputArray[last]; i++ {
		if !isContain(&inputArray, i) {
			if isValid(&inputArray, i) {
				return i
			}
		}
	}
	return inputArray[last] + 1
}

func boxBlur(image [][]int) [][]int {

	res := [][]int{}
	for row := 2; row < len(image); row++ {
		line := []int{}
		for col := 2; col < len(image[0]); col++ {
			fmt.Printf("row: %d, col: %d\n", row, col)
			counter := 0
			for i := row; i >= row - 2; i-- {
				for j := col; j >= col - 2; j-- {
					counter += image[i][j]
				}
			}
			fmt.Print(" ", counter / 9)
			line = append(line, counter / 9)
		}
		res = append(res, line)
		fmt.Println()

	}
	return res

}

func initArray(row, col int) [][]int {
	res := [][]int{}
	for i := 0; i < row; i++ {
		line := []int{}
		for j := 0; j < col; j++ {
			line = append(line, 0)
		}
		res = append(res, line)
	}
	return res
}
func updateMinesMatrix(i, j int, matrix *[][]int) {
	for k := i - 1; k < i + 2; k++ {
		for t := j - 1; t < j + 2; t++ {
			if k == i && t == j {
				continue
			}
			if (k >= 0 && k < len(*matrix)) && (t >= 0 && t < len((*matrix)[0])) {
				(*matrix)[k][t] += 1
			}
		}
	}
}

func minesweeper(matrix [][]bool) [][]int {
	res := initArray(len(matrix), len(matrix[0]))

	for i, line := range matrix {
		for j, el := range line {
			if el {
				updateMinesMatrix(i, j, &res)
			}
		}
	}
	return res
}
