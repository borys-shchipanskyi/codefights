package intro

func growingPlant(upSpeed int, downSpeed int, desiredHeight int) int {
	level := 0
	days := 0
	for true {
		days += 1
		level += upSpeed
		if level >= desiredHeight{
			return days
		}
		level -= downSpeed
	}
	return 0

}

func knapsackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
	if weight1 + weight2 <= maxW{
		return value1 + value2
	}
	if weight1 <= maxW && weight2 <= maxW{
		if value1 >= value2{
			return value1
		}
		return value2
	}
	if weight2 <= maxW{
		return value2
	}
	if weight1 <= maxW {
		return value1
	}
	return 0
}


func longestDigitsPrefix(inputString string) string {
	res := ""
	for i := 0; i < len(inputString); i++ {
		cr := inputString[i]
		if cr > '9' || cr < '0'{
			break
		}
		res +=string(cr)

	}
	return res
}

func getNumSum(n int)int {
	if n < 10{
		return n
	}
	return n%10 + getNumSum(n / 10)
}

func digitDegree(n int) int {
	if n < 10{
		return 0
	}
	counter := 1
	res := getNumSum(n)
	for res >= 10 {
		res = getNumSum(res)
		counter += 1
	}
	return counter
}
//
//func getSelIndexes(cell string) (int, int){
//	row, _ := strconv.Atoi(cell[1:])
//	col := int(cell[0] - 65)
//	return row -1, col
//}

func bishopAndPawn(bishop string, pawn string) bool {
	bRow, bCel := getSelIndexes(bishop)
	pRow, pCel := getSelIndexes(pawn)

	step := 0
	if bRow > pRow{
		step = bRow - pRow
	} else {
		step = pRow - bRow
	}

	if bCel > pCel{
		return bCel - step == pCel
	}
	return bCel + step == pCel

}
