package tests

import (
	. "codefights/interviewPractice/solutions"
	"testing"

	"fmt"
)



func TestGroupingDishes(t *testing.T) {
	tests := []struct {
		input  [][]string
		out [][]string
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
		fmt.Println("************************")
		PrintHashTable(test.input)
		fmt.Println("Got res: ")
		PrintHashTable(res)
		if !IsHashTablesEquals(test.out, res) {
			t.Errorf("Test input: %#v, Expect: %#v, but get: %#v\n", test.input, test.out, res)
		}
	}

}
