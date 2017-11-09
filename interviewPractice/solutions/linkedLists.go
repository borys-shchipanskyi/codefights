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
	if (currentL2 != nil) {
		return false
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

func getSumAndOdd(sum int) (int, int) {
	n := 10000
	if sum >= n {
		return sum % n, sum / n
	}
	return sum, 0

}

func AddTwoHugeNumbers(a *ListNode, b *ListNode) *ListNode {
	var res *ListNode = nil
	var first *ListNode = nil
	revA := ReverseList(a)
	revB := ReverseList(b)

	for odd := 0; revB != nil || revA != nil || odd != 0; {
		sum := 0
		if revB != nil && revA != nil {
			sum, odd = getSumAndOdd(revB.Value.(int) + revA.Value.(int) + odd)
			revB = revB.Next
			revA = revA.Next

		} else if revB == nil && revA != nil {
			sum, odd = getSumAndOdd(revA.Value.(int) + odd)
			revA = revA.Next
		} else if revA == nil && revB != nil {
			sum, odd = getSumAndOdd(revB.Value.(int) + odd)
			revB = revB.Next
		} else {
			sum, odd = getSumAndOdd(odd)
		}

		if res == nil {
			res = &ListNode{sum, nil}
			first = res
		} else {
			res.Next = &ListNode{sum, nil}
			res = res.Next
		}
	}

	return ReverseList(first)
}

func MergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode = nil
	var head *ListNode = nil
	for l1 != nil || l2 != nil {
		val := 0
		if l1 != nil && l2 != nil {
			if l1.Value.(int) < l2.Value.(int) {
				val = l1.Value.(int)
				l1 = l1.Next
			} else {
				val = l2.Value.(int)
				l2 = l2.Next
			}
		} else if l1 == nil {
			val = l2.Value.(int)
			l2 = l2.Next
		} else if l2 == nil {
			val = l1.Value.(int)
			l1 = l1.Next
		}

		if res == nil {
			res = &ListNode{val, nil}
			head = res
		} else {
			res.Next = &ListNode{val, nil}
			res = res.Next
		}
	}
	return head

}

func SplitList(l *ListNode, k int) (head *ListNode, tail *ListNode) {
	head = l
	var subListTail *ListNode
	for i, current := 0, l; i < k && current != nil; i++ {
		//fmt.Printf("i: %d \n", i)
		subListTail = current
		current = current.Next
	}
	if subListTail == nil {
		return nil, nil
	}
	tail = subListTail.Next
	subListTail.Next = nil

	return head, tail
}

func getListLen(l *ListNode)int{
	counter := 0
	for current:= l; current != nil; current = current.Next{
		counter += 1
	}
	return counter
}

func ReverseNodesInKGroups(l *ListNode, k int) *ListNode {
	if l.Next == nil{
		return l
	}
	var head, tail, buff, newTail *ListNode
	counter := 0
	for current := l; current != nil; {
		if counter == k {
			counter = 0
			if head == nil {
				head = buff
				tail = newTail
			} else {
				tail.Next = buff
				tail = newTail

			}
			newTail = nil
		}
		counter += 1
		if newTail == nil {
			newTail = current
			buff = newTail
			current = current.Next
			newTail.Next = nil
		} else {
			tmp := current
			current = current.Next
			tmp.Next = buff
			buff = tmp
		}

	}
	if getListLen(buff) != k{
		tail.Next = ReverseList(buff)
	} else {
		tail.Next = buff
	}

	return head
}