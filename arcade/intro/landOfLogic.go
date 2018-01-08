package intro

import (
	"strings"
	"strconv"
	"fmt"
)

//
func isLetter(cr byte) bool {
	return (cr >= 'a' && cr <= 'z') || (cr >= 'A' && cr <= 'Z')
}

func longestWord(text string) string {
	max := ""
	for i := 0; i < len(text); i++ {
		tmp := ""

		for j := i; j < len(text); j++ {
			cr := text[j]
			if isLetter(cr) {
				tmp += string(cr)
			} else {
				i = j
				break
			}
		}
		if len(tmp) > 0 {
			if len(max) < len(tmp) {
				max = tmp
			}
		}
	}
	return max
}

//

func validTime(time string) bool {
	res := strings.Split(time, ":")
	h, _ := strconv.Atoi(res[0])
	m, _ := strconv.Atoi(res[1])
	return (h >= 0 && h < 24) && (m >= 0 && m < 60)
}

//
func sumUpNumbers(inputString string) int {
	numbers := []int{}
	for i := 0; i < len(inputString); i++ {
		tmp := ""
		j := i
		for ; j < len(inputString); j++ {
			cr := inputString[j]
			if cr >= '0' && cr <= '9' {
				tmp += string(cr)
			} else {

				break
			}
		}
		i = j
		if len(tmp) > 0 {
			n, _ := strconv.Atoi(tmp)
			numbers = append(numbers, n)
		}
	}
	fmt.Println(inputString, numbers)
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

//
func differentSquares(matrix [][]int) int {

	res := map[string]bool{}
	for i := 0; i < len(matrix) - 1; i++ {
		for j := 0; j < len(matrix[0]) - 1; j++ {
			s := strconv.Itoa(matrix[i][j]) + strconv.Itoa(matrix[i + 1][j]) + strconv.Itoa(matrix[i][j + 1]) + strconv.Itoa(matrix[i + 1][j + 1])
			res[s] = true
		}
	}
	return len(res)
}
//
func gerDiv(n int) int {
	for i := 9; i > 1; i-- {
		if n % i == 0 {
			return i
		}
	}
	return -1
}

func digitsProduct(product int) int {
	if product == 0 {
		return 10
	}
	res := 0
	mv := 1
	for product >= 10 {
		div := gerDiv(product)
		if div == -1 {
			return div
		}
		res = div * mv + res
		mv *= 10
		product /= div
	}

	res = product * mv + res
	return res
}
//
func fileNaming(names []string) []string {
	m := map[string]int{}
	res := []string{}
	for _, s := range names {
		m[s] += 1
		if (m[s] > 1) {
			for z := 1; ; z++ {
				dup := s + "(" + strconv.Itoa(z) + ")"
				if m[dup] == 0 {
					m[dup] += 1
					res = append(res, dup)
					break
				}
			}
		} else {
			res = append(res, s)
		}
	}

	return res
}

//
func pow2(pow int) int {
	res := 1
	for i := 1; i <= pow; i++ {
		res *= 2
	}
	return res
}
func binToDec(s string) byte {
	res := byte(0)
	for i := 0; i < len(s); i++ {
		el := len(s) - 1 - i
		if s[el] == '1' {
			res += byte(pow2(i))
		}

	}
	return res

}
func messageFromBinaryCode(code string) string {
	res := []byte{}
	for i := 8; i <= len(code); i += 8 {
		l := code[i - 8: i]
		res = append(res, binToDec(l))
	}
	return string(res)
}

//
func spiralNumbers(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	i, j := 0, 0
	mv := 0
	target := n * n
	for k := 1; k <= target; {

		for j = mv; j < n - mv && k <= target; j++ {
			res[i][j] = k
			k++
		}
		j--
		for i = mv + 1; i < n - mv && k <= target; i++ {
			res[i][j] = k
			k++
		}
		i--
		for j = n - mv - 2; j >= mv && k <= target; j-- {
			res[i][j] = k
			k++
		}

		j = mv
		for i = n - mv - 2; i > mv && k <= target; i-- {
			res[i][j] = k
			k++
		}
		i++
		mv++

	}
	return res
}

//
func sudoku(grid [][]int) bool {

	for i := 0; i < len(grid); i++ {
		row := map[int]bool{}
		col := map[int]bool{}
		for j := 0; j < len(grid); j++ {
			rowVal := grid[i][j]
			colVal := grid[j][i]
			if _, ok := row[rowVal]; ok {
				return false
			} else {
				row[rowVal] = true
			}
			if _, ok := col[colVal]; ok {
				return false
			} else {
				col[colVal] = true
			}
		}
	}
	end := 3
	for i := 0; i < 9; i++ {
		values := map[int]bool{}
		pos := i - i % 3
		for k := pos; k < pos + 3; k++ {
			for j := end - 3; j < end; j++ {
				if _, ok := values[grid[k][j]]; ok {
					return false
				} else {
					values[grid[k][j]] = true
				}
			}
		}
		end += 3
		if (end > 9) {
			end = 3
		}
	}

	return true
}
