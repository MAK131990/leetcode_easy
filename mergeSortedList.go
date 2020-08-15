package main

// func main() {
// 	var a *ListNode
// 	fmt.Println(a)
// }

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head *ListNode

	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

	tail := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	for l1 != nil {
		tail.Next = l1
		l1 = l1.Next
		tail = tail.Next
	}
	for l2 != nil {
		tail.Next = l2
		l2 = l2.Next
		tail = tail.Next
	}

	return head
}
