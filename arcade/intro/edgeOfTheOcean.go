package intro

func adjacentElementsProduct(inputArray []int) int {
	max := -1001
	for i := 0; i < len(inputArray) - 1; i++ {
		product := inputArray[i] * inputArray[i + 1]
		if max < product {
			max = product
		}

	}
	return max
}

func shapeArea(n int) int {
	return n * n + (n - 1) * (n - 1)
}

func makeArrayConsecutive2(statues []int) int {
	//sort array
	last := len(statues) - 1
	for i := 0; i <= last; i++ {
		for j := i + 1; j <= last; j++ {
			if statues[i] > statues[j] {
				tmp := statues[i]
				statues[i] = statues[j]
				statues[j] = tmp
			}
		}
	}
	return statues[last] - statues[0] - last
}

func almostIncreasingSequence(x []int) bool {
	min := x[0]
	max := x[len(x) - 1]
	minN := 0
	maxN := 0
	for i := 1; i < len(x); i++ {
		// j-=1
		if x[i] <= min {
			minN += 1
		} else {
			min = x[i]
		}
	}

	for i := len(x) - 2; i >= 0; i-- {
		if x[i] >= max {
			maxN += 1
		} else {
			max = x[i]
		}
	}

	return !(minN > 1 && maxN > 1)
}

func isContain(usedIndexes []int, target int) bool {
	for _, el := range usedIndexes {
		if el == target {
			return true
		}
	}
	return false
}

func matrixElementsSum(matrix [][]int) int {
	usedIndexes := []int{}
	sum := 0
	for i, _ := range matrix {
		for j, el := range matrix[i] {
			if el == 0 {
				usedIndexes = append(usedIndexes, j)
			} else {
				if !isContain(usedIndexes, j) {
					sum += el
				}
			}

		}
	}
	return sum
}
