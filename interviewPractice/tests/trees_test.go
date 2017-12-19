package tests

import (
	. "codefights/interviewPractice/solutions"
	"testing"
)

func TestHasPathWithGivenSum(t *testing.T) {
	tests := []struct {
		input *Tree
		s     int
		out   bool
	}{
		{
			CreateTree(StrToMap([]byte(`{"value": 4,"left": {"value": 1,"left": {"value": -2,"left": null,
			"right": {"value": 3,"left": null,"right": null}},"right": null},"right": {"value": 3,"left":
			{"value": 1,"left": null,"right": null},"right": {"value": 2,"left": {"value": -2,"left": null,
			"right": null},"right": {"value": -3,"left": null,"right": null}}}}`))),
			7,
			true,
		},
		{
			CreateTree(StrToMap([]byte(`{"value": 4,"left": {"value": 1,"left": {"value": -2,"left": null,
			"right": {"value": 3,"left": null,"right": null}},"right": null},"right": {"value": 3,"left":
			{"value": 1,"left": null,"right": null},"right": {"value": 2,"left": {"value": -4,"left": null,
			"right": null},"right": {"value": -3,"left": null,"right": null}}}}`))),
			7,
			false,
		},
		{
			CreateTree(StrToMap([]byte(`{"value": 8,"left": null,"right": {"value": 3,"left": null,"right": null}}`))),
			8,
			false,
		},
	}
	for _, test := range tests {

		res := HasPathWithGivenSum(test.input, test.s)
		if res != test.out {
			t.Errorf("Test input: %#v, Expect: %#v, but get: %#v\n", test.input, test.out, res)
		}
	}
}

func TestFindProfession(t *testing.T){
	tests := []struct {
		level int
		pos     int
		out   string
	}{
		{
			3,
			3,
			"Doctor",
		},

	}

	for _, test := range tests {

		res := FindProfession(test.level, test.pos)
		if res != test.out {
			t.Errorf("Test input: %#v, Expect: %#v, but get: %#v\n", test.level, test.out, res)
		}
	}
}

func TestKthSmallestInBST(t *testing.T) {
	tests := []struct {
		input *Tree
		s     int
		out   int
	}{
		{
			CreateTree(StrToMap([]byte(`{"value": 3,"left": {"value": 1,"left": null,"right": null},
			"right": {"value": 5,"left": {"value": 4,"left": null,"right": null},"right": {"value": 6,
			"left": null,"right": null}}}`))),
			4,
			5,
		},

	}
	for _, test := range tests {

		res := KthSmallestInBST(test.input, test.s)
		if res != test.out {
			t.Errorf("Test input: %#v, Expect: %#v, but get: %#v\n", test.input, test.out, res)
		}
	}
}

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		t1  *Tree
		t2  *Tree
		out bool
	}{
		{
			CreateTree(StrToMap([]byte(`{
    "value": 5,
    "left": {
        "value": 10,
        "left": {
            "value": 4,
            "left": {
                "value": 1,
                "left": null,
                "right": null
            },
            "right": {
                "value": 2,
                "left": null,
                "right": null
            }
        },
        "right": {
            "value": 6,
            "left": null,
            "right": {
                "value": -1,
                "left": null,
                "right": null
            }
        }
    },
    "right": {
        "value": 7,
        "left": null,
        "right": null
    }
}`))),
			CreateTree(StrToMap([]byte(`{
    "value": 10,
    "left": {
        "value": 4,
        "left": {
            "value": 1,
            "left": null,
            "right": null
        },
        "right": {
            "value": 2,
            "left": null,
            "right": null
        }
    },
    "right": {
        "value": 6,
        "left": null,
        "right": {
            "value": -1,
            "left": null,
            "right": null
        }
    }
}`))),
			true,
		},

	}
	for _, test := range tests {

		res := IsSubtree(test.t1, test.t2)
		if res != test.out {
			t.Errorf("Test input: %#v, Expect: %#v, but get: %#v\n", test.t1, test.out, res)
		}
	}
}