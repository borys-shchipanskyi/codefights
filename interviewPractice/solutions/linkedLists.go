package solutions

import "fmt"

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func PrintListNode(l *ListNode) {
	current := l
	for (current != nil) {
		fmt.Printf("%d ", current.Value)
		current = current.Next
	}
	fmt.Println()
}

func IsEqual(l1 *ListNode, l2 *ListNode) bool {
	currentL1 := l1
	currentL2 := l2
	for (currentL1 != nil ) {
		if (currentL2 == nil) {
			return false
		}
		if currentL1.Value != currentL2.Value {
			return false
		}
		currentL1 = currentL1.Next
		currentL2 = currentL2.Next
	}
	return true
}

func RemoveKFromList(l *ListNode, k int) *ListNode {
	if (l == nil) {
		return l
	}
	l.Next = RemoveKFromList(l.Next, k)
	if (l.Value == k) {
		return l.Next
	}
	return l
}

func ReverseList(l *ListNode) *ListNode {
	current := l.Next
	end := &ListNode{l.Value, nil}
	for (current != nil) {
		tmp := current
		current = current.Next
		tmp.Next = end
		end = tmp
	}
	return end
}

func IsListPalindrome(l *ListNode) bool {
	isPalindrome := true
	if l == nil {
		return isPalindrome
	}
	middleNode := l

	for current := l; current != nil && current.Next != nil; current = current.Next {
		middleNode = middleNode.Next
		current = current.Next

	}

	for headCounter, tailCounter := l, ReverseList(middleNode); tailCounter != nil; {
		if (headCounter.Value != tailCounter.Value) {
			isPalindrome = false
			break
		}
		headCounter = headCounter.Next
		tailCounter = tailCounter.Next
	}

	return isPalindrome

}