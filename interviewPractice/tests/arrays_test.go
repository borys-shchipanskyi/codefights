package tests

import (
	. "codefights/interviewPractice/solutions"
	"reflect"
	"testing"
)

func TestFirstDuplicate(t *testing.T) {
	tests := []struct {
		input []int
		out   int
	}{
		{[]int{2, 3, 3, 1, 5, 2}, 3},
		{[]int{2, 4, 3, 5, 1}, -1},
		{[]int{1}, -1},
		{[]int{2, 2}, 2},
		{[]int{2, 1}, -1},
		{[]int{2, 1, 3}, -1},
		{[]int{2, 3, 3}, 3},
		{[]int{3, 3, 3}, 3},
		{[]int{8, 4, 6, 2, 6, 4, 7, 9, 5, 8}, 6},
		{[]int{10, 6, 8, 4, 9, 1, 7, 2, 5, 3}, -1},
		{[]int{1, 1, 2, 2, 1}, 1},
	}

	for _, test := range tests {

		res := FirstDuplicate(test.input)
		if res != test.out {
			t.Errorf("Test input: %#v, Expect: %d, but get: %d\n", test.input, test.out, res)
		}
	}
}

func TestFirstNotRepeatingCharacter(t *testing.T) {
	tests := []struct {
		input string
		out   string
	}{
		{"abacabad", "c"},
		{"abacabaabacaba", "_"},
		{"z", "z"},
		{"bcb", "c"},
		{"bcccccccb", "_"},
		{"abcdefghijklmnopqrstuvwxyziflskecznslkjfabe", "d"},
		{"zzz", "_"},
		{"bcccccccccccccyb", "y"},
		{"xdnxxlvupzuwgigeqjggosgljuhliybkjpibyatofcjbfxwtalc", "d"},
		{"ngrhhqbhnsipkcoqjyviikvxbxyphsnjpdxkhtadltsuxbfbrkof", "g"},
	}

	for _, test := range tests {

		res := FirstNotRepeatingCharacter(test.input)
		if res != test.out {
			t.Errorf("Test input: %#v, Expect: %s, but get: %s\n", test.input, test.out, res)
		}
	}

}

func TestRotateImage(t *testing.T) {
	tests := []struct {
		input [][]int
		out   [][]int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			[][]int{{1}},
			[][]int{{1}},
		},
		{
			[][]int{{10, 9, 6, 3, 7}, {6, 10, 2, 9, 7}, {7, 6, 3, 8, 2}, {8, 9, 7, 9, 9}, {6, 8, 6, 8, 2}},
			[][]int{{6, 8, 7, 6, 10}, {8, 9, 6, 10, 9}, {6, 7, 3, 2, 6}, {8, 9, 8, 9, 3}, {2, 9, 2, 7, 7}},
		},
		{
			[][]int{{40, 12, 15, 37, 33, 11, 45, 13, 25, 3},
				{37, 35, 15, 43, 23, 12, 22, 29, 46, 43},
				{44, 19, 15, 12, 30, 2, 45, 7, 47, 6},
				{48, 4, 40, 10, 16, 22, 18, 36, 27, 48},
				{45, 17, 36, 28, 47, 46, 8, 4, 17, 3},
				{14, 9, 33, 1, 6, 31, 7, 38, 25, 17},
				{31, 9, 17, 11, 29, 42, 38, 10, 48, 6},
				{12, 13, 42, 3, 47, 24, 28, 22, 3, 47},
				{38, 23, 26, 3, 23, 27, 14, 40, 15, 22},
				{8, 46, 20, 21, 35, 4, 36, 18, 32, 3}},
			[][]int{{8, 38, 12, 31, 14, 45, 48, 44, 37, 40},
				{46, 23, 13, 9, 9, 17, 4, 19, 35, 12},
				{20, 26, 42, 17, 33, 36, 40, 15, 15, 15},
				{21, 3, 3, 11, 1, 28, 10, 12, 43, 37},
				{35, 23, 47, 29, 6, 47, 16, 30, 23, 33},
				{4, 27, 24, 42, 31, 46, 22, 2, 12, 11},
				{36, 14, 28, 38, 7, 8, 18, 45, 22, 45},
				{18, 40, 22, 10, 38, 4, 36, 7, 29, 13},
				{32, 15, 3, 48, 25, 17, 27, 47, 46, 25},
				{3, 22, 47, 6, 17, 3, 48, 6, 43, 3}},
		},
	}

	for _, test := range tests {

		res := RotateImage(test.input)
		if !reflect.DeepEqual(res, test.out) {
			t.Errorf("Test input: %#v, Expect: %s, but get: %s\n", test.input, test.out, res)
		}
	}
}

func TestSudoku2(t *testing.T) {
	tests := []struct {
		input [][]string
		out   bool
	}{
		{[][]string{{".", ".", ".", "1", "4", ".", ".", "2", "."},
			{".", ".", "6", ".", ".", ".", ".", ".", "."},
			{".", ".", ".", ".", ".", ".", ".", ".", "."},
			{".", ".", "1", ".", ".", ".", ".", ".", "."},
			{".", "6", "7", ".", ".", ".", ".", ".", "9"},
			{".", ".", ".", ".", ".", ".", "8", "1", "."},
			{".", "3", ".", ".", ".", ".", ".", ".", "6"},
			{".", ".", ".", ".", ".", "7", ".", ".", "."},
			{".", ".", ".", "5", ".", ".", ".", "7", "."}}, true},
	}
	for _, test := range tests {

		res := Sudoku2(test.input)
		if res != test.out {
			t.Errorf("Test input: %#v, Expect: %s, but get: %s\n", test.input, test.out, res)
		}
	}
}


func TestIsCryptSolution(t *testing.T) {
	tests := []struct {
		crypt    []string
		solution [][]string
		out      bool
	}{
		{
			[]string{"SEND", "MORE", "MONEY"},
			[][]string{{"O", "0"}, {"M", "1"}, {"Y", "2"}, {"E", "5"}, {"N", "6"}, {"D", "7"}, {"R", "8"}, {"S", "9"}},
			true,
		},

	}
	for _, test := range tests {

		res := IsCryptSolution(test.crypt, test.solution)
		if res != test.out {
			t.Errorf("Test input: {crypt: %#v, solution: %#v}, Expect: %s, but get: %s\n",
				test.crypt, test.solution, test.out, res)
		}
	}
}
