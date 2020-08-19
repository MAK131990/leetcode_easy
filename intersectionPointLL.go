package main

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	headALen, headBLen, temp := 0, 0, headA
	for temp.Next != nil {
		headALen++
		temp = temp.Next
	}
	for temp = headB; temp.Next != nil; {
		headBLen++
		temp = temp.Next
	}
	startA, startB := headA, headB
	for headALen < headBLen && startA != nil {
		startA = startA.Next
		headALen++
	}
	for headBLen < headALen && startB != nil {
		startB = startB.Next
		headBLen++
	}
	if startA == nil || startB == nil {
		return nil
	}
	for startA != nil && startB != nil {
		if startA == startB {
			return startA
		}
		startA = startA.Next
		startB = startB.Next
	}
	return nil
}
