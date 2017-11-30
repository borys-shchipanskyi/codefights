package tests

import (
	. "codefights/interviewPractice/solutions"
	"testing"

	"fmt"
)

func TestGroupingDishes(t *testing.T) {
	tests := []struct {
		input [][]string
		out   [][]string
	}{
		{
			[][]string{{"Salad", "Tomato", "Cucumber", "Salad", "Sauce"},
				{"Pizza", "Tomato", "Sausage", "Sauce", "Dough"},
				{"Quesadilla", "Chicken", "Cheese", "Sauce"},
				{"Sandwich", "Salad", "Bread", "Tomato", "Cheese"}},
			[][]string{{"Cheese", "Quesadilla", "Sandwich"},
				{"Salad", "Salad", "Sandwich"},
				{"Sauce", "Pizza", "Quesadilla", "Salad"},
				{"Tomato", "Pizza", "Salad", "Sandwich"}},
		},
	}
	for _, test := range tests {
		res := GroupingDishes(test.input)
		PrintHashTable(test.input)
		fmt.Println("Got res: ")
		PrintHashTable(res)
		if !IsHashTablesEquals(test.out, res) {
			t.Errorf("Test input: %#v, Expect: %#v, but get: %#v\n", test.input, test.out, res)
		}
	}

}

func TestAreFollowingPatterns(t *testing.T) {
	tests := []struct {
		strings  []string
		patterns []string
		out      bool
	}{
		{
			[]string{"cat", "dog", "dog"},
			[]string{"a", "b", "b"},
			true,
		},
		{
			[]string{"cat", "dog", "dog"},
			[]string{"a", "b", "c"},
			false,
		},
		{
			[]string{"cat", "dog", "doggy"},
			[]string{"a", "b", "b"},
			false,
		},

	}

	for _, test := range tests {
		if AreFollowingPatterns(test.strings, test.patterns) != test.out {
			t.Error("Test input error")
			fmt.Printf("%#v\n", test)
		}
	}
}

func TestContainsCloseNums(t *testing.T) {
	tests := []struct {
		coins    []int
		quantity []int
		out      int
	}{
		{
			[]int{10, 50, 100},
			[]int{1, 2, 1},
			9,
		},
		{
			[]int{10, 50, 100, 500},
			[]int{5, 3, 2, 2},
			122,
		},
		{
			[]int{1, 2},
			[]int{50000, 2},
			50004,
		},
	}

	for _, test := range tests {
		res := PossibleSums(test.coins, test.quantity)
		if res != test.out {
			t.Error("Test input error")
			fmt.Printf("%#v\n", res)
		}
	}
}

func TestPossibleSums(t *testing.T) {
	tests := []struct {
		coins    []int
		quantity []int
		out      int
	}{
		{
			[]int{10, 50, 100},
			[]int{1, 2, 1},
			9,
		},
		{
			[]int{10, 50, 100, 500},
			[]int{5, 3, 2, 2},
			122,
		},
	}

	for _, test := range tests {
		res := PossibleSums(test.coins, test.quantity)
		if res != test.out {
			t.Error("Test input error")
		}
	}

}

func TestSwapLexOrder(t *testing.T) {
	tests := []struct {
		str   string
		pairs [][]int
		out   string
	}{
		{
			"abdc",
			[][]int{{1, 4}, {3, 4}},
			"dbca",
		},
		{
			"acxrabdz",
			[][]int{{1, 3}, {6, 8}, {3, 8}, {2, 7}},
			"zdxrabca",
		},
	}

	for _, test := range tests {
		res := SwapLexOrder(test.str, test.pairs)
		if res != test.out {
			t.Error("Test input error")
		}
	}

}