package main

// func main() {

// }

//ListNodeLL definition of ListNode of LInkList
type ListNodeLL struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	temp := head
	for temp.Next != nil {
		if temp.Val != temp.Next.Val {
			temp = temp.Next
		} else {
			temp.Next = temp.Next.Next
		}
	}
	return head
}
