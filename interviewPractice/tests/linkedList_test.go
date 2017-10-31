package tests

import (
	. "codefights/interviewPractice/solutions"
	"testing"
	"fmt"
)

func TestRemoveKFromList(t *testing.T) {
	tests := []struct {
		input *ListNode
		k     int
		out   *ListNode
	}{
		{
			&ListNode{3, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}},
			3,
			&ListNode{1, &ListNode{2, &ListNode{4, &ListNode{5, nil}}}},
		},
		{
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}},
			10,
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, nil}}}}}}},
		},
		{
			&ListNode{1000, &ListNode{1000, nil}},
			1000,
			nil,
		},
		{
			nil,
			-1000,
			nil,
		},
		{
			&ListNode{123, &ListNode{456, &ListNode{789, &ListNode{0, nil}}}},
			0,
			&ListNode{123, &ListNode{456, &ListNode{789, nil}}},
		},
	}

	for _, test := range tests {

		res := RemoveKFromList(test.input, test.k)
		if !IsEqual(test.out, res) {
			t.Errorf("Test input: %#v, Expect: %#v, but get: %#v\n", test.input, test.out, res)
		}
	}
}

func TestIsListPalindrome(t *testing.T) {
	tests := []struct {
		input *ListNode
		out   bool
	}{
		{&ListNode{0, &ListNode{1, &ListNode{0, nil}}}, true},
		{&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, nil}}}}, false},
		{&ListNode{1, &ListNode{-1000000000, &ListNode{-1000000000, &ListNode{1000000000, &ListNode{1, nil}}}}}, false},
	}
	for _, test := range tests {
		//PrintListNode(test.input)
		res := IsListPalindrome(test.input)
		if res != test.out {

			PrintListNode(test.input)
			t.Errorf("Test input: %#v, Expect: %#v, but get: %#v\n", test.input, test.out, res)
		}
	}
}

func TestAddTwoHugeNumbers(t *testing.T) {
	tests := []struct {
		a *ListNode
		b *ListNode
		out   *ListNode
	}{
		{
			&ListNode{9876, &ListNode{5432, &ListNode{1999, nil}}},
			&ListNode{1, &ListNode{8001, nil}},
			&ListNode{9876, &ListNode{5434, &ListNode{0, nil}}},
		},
		{
			&ListNode{1, nil},
			&ListNode{9998, &ListNode{9999, &ListNode{9999, &ListNode{9999, &ListNode{9999, &ListNode{9999, nil}}}}}},
			&ListNode{9999, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, nil}}}}}},
		},
		{
			&ListNode{1, nil},
			&ListNode{9999, &ListNode{9999, &ListNode{9999, &ListNode{9999, &ListNode{9999, &ListNode{9999, nil}}}}}},
			&ListNode{1, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{0, nil}}}}}}},
		},

	}
	for _, test := range tests {
		res := AddTwoHugeNumbers(test.a, test.b)
		PrintListNode(res)
		fmt.Printf("res : %v\n", IsEqual(res, test.out))
		if !IsEqual(res, test.out) {
			t.Errorf("Test input: %#v, Expect: %#v, but get: %#v\n", test.a, test.out, res)
		}
	}
}


func TestMergeTwoLinkedLists(t *testing.T) {
	tests := []struct {
		l1 *ListNode
		l2 *ListNode
		out   *ListNode
	}{
		{
			&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
			&ListNode{4, &ListNode{5, &ListNode{6, nil}}},
			&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
		},

	}
	for _, test := range tests {
		res := MergeTwoLinkedLists(test.l1, test.l2)
		PrintListNode(res)
		fmt.Printf("res : %v\n", IsEqual(res, test.out))
		if !IsEqual(res, test.out) {
			t.Errorf("Test input: %#v, Expect: %#v, but get: %#v\n", test.l1, test.out, res)
		}
	}
}
