package intro

import (
	"strconv"
	"sort"
)

func isDigit(symbol string) bool {
	if len(symbol) > 1{
		return false
	}
	return symbol[0] >='0' && symbol[0] <='9'
}

func lineEncoding(s string) string {
	res := ""
	for i:= 0; i < len(s); i++{
		cr := s[i]
		count := 0
		j := i
		for ; j<len(s); j++{
			if s[j] != cr{
				break
			}
			count += 1
		}
		i = j -1
		if count > 1{
			res += strconv.Itoa(count)
		}
		res += string(cr)
	}
	return res

}


func getCelIndexes(cell string) (int, int){
	row, _ := strconv.Atoi(cell[1:])
	col := int(cell[0] - 97)
	return row, col +1
}

func isValidPosition(row, col int)bool{
	return row <=8 && row >0 && col <= 8 && col > 0
}

func chessKnight(cell string) int {
	posShift := []int{-2,-1,1,2}
	counter := 0
	row, col := getCelIndexes(cell)
	for _, el := range posShift{
		tmp := []int{-2,2}
		if el == -2 || el == 2{
			tmp = []int{-1,1}
		}
		for _, e := range tmp{
			if isValidPosition(row + el, col + e){
				counter +=1
			}
		}
	}
	return counter

}

func deleteDigit(n int) int {
	res := []int{}
	last := 0
	shift := 1
	for n > 0{
		tmp := n % 10
		n /=10
		val := n * shift + last
		res = append(res, val)
		last = tmp * shift + last
		shift *= 10
	}
	sort.Ints(res)

	return res[len(res) -1]

}
