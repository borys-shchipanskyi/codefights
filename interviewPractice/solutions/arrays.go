package solutions

import "strconv"

func FirstDuplicate(a []int) int {
	for _, val := range a {
		if val < 0 {
			if a[-1 * val -1] < 0{
				return -1 * val
			}
			a[-1 * val -1]*= -1

		}else{
			if a[val -1] < 0{
				return val
			}
			a[val - 1] *= -1
		}

	}
	return -1;

}

func FirstNotRepeatingCharacter(s string) string {
	symbols := []string{}
	count := []int{}
	for _, sb := range s {
		el := string(sb)
		poss := -1
		for index, _ := range symbols {
			if (symbols[index] == el) {
				poss = index
				count[poss] += 1
				break;
			}
		}
		if (poss == -1) {
			symbols = append(symbols, el)
			count = append(count, 1)
		}
	}
	for index, val := range count {
		if (val == 1) {
			return symbols[index]
		}
	}

	return "_"
}

func RotateImage(a [][]int) [][]int {
	length := len(a)
	var b = make([][]int, len(a))
	for i := 0; i < length; i++{
		for j := length -1; j >=0; j--{
			b[i] = append(b[i], a[j][i])
		}
	}
	return b
}

func Sudoku2(grid [][]string) bool {
	for i, _ := range grid[0] {
		vertical := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
		horisontal := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
		for j, _ := range grid[0] {
			if (grid[i][j] != ".") {
				el, _ := strconv.Atoi(grid[i][j])
				if vertical[el] != -1 {
					return false
				}
				vertical[el] = el
			}
			if (grid[j][i] != ".") {
				el, _ := strconv.Atoi(grid[j][i])
				if horisontal[el] != -1 {
					return false
				}
				horisontal[el] = el
			}
		}
	}
	end := 3
	for i := 0; i < 9; i++ {
		values := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
		pos := i - i % 3
		for k := pos; k < pos + 3; k++ {
			for j := end - 3; j < end; j++ {
				if (grid[k][j] != ".") {
					el, _ := strconv.Atoi(grid[k][j])
					if values[el] != -1 {
						return false
					}
					values[el] = el
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

// isCryptSolution
func getNumberFromSolutionByLetter(letter string, solution *[][]string)string{
	for _, line := range *solution{
		if (line[0] == letter){
			return line[1]
		}
	}
	return ""
}
func cryptWord(word string, solution *[][]string)string{
	var crypt = ""
	for  i := 0; i < len(word); i++ {
		crypt += getNumberFromSolutionByLetter(string(word[i]), solution)
	}
	return crypt
}

func IsCryptSolution(crypt []string, solution [][]string) bool {
	w1 :=  cryptWord(crypt[0], &solution)
	w2 :=  cryptWord(crypt[1], &solution)
	w3 :=  cryptWord(crypt[2], &solution)
	if ((w1[0] == '0' && len(w1) > 1) || (w2[0] == '0' && len(w2) > 1 ) || (w3[0] == '0' && len(w2) > 1)){
		return false
	}
	firstNumber, _ := strconv.Atoi(w1)
	secondNumber, _ := strconv.Atoi(w2)
	result, _ := strconv.Atoi(w3)
	return firstNumber + secondNumber == result
}