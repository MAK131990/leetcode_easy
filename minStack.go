package main

import "fmt"

// [,"top","getMin","push","top","getMin","pop","getMin"]
// [,[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]
func main() {
	obj := Constructor()
	obj.Push(2147483646)
	obj.Push(2147483646)
	obj.Push(2147483647)
	fmt.Println(obj.Top())
	obj.Pop()
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
	obj.Pop()
	obj.Push(2147483647)
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
	obj.Push(-2147483648)
	fmt.Println(obj.Top())
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())
}

//MinStack struct
type MinStack struct {
	top      int
	stack    []int
	minStack []int
}

//Constructor ** initialize your data structure here. */
func Constructor() MinStack {
	var obj MinStack
	return obj
}

//Push function for stacj
func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if this.top == 0 || (this.top > 0 && this.GetMin() > x) {
		this.minStack = append(this.minStack, x)
	} else {
		this.minStack = append(this.minStack, this.GetMin())
	}
	this.top++
}

//Pop function for stack
func (this *MinStack) Pop() {
	if this.top > 0 {
		this.minStack = this.minStack[0 : this.top-1]
		this.stack = this.stack[0 : this.top-1]
		this.top--
	}
}

//Top returns Top element in stack
func (this *MinStack) Top() int {
	return this.stack[this.top-1]
}

//GetMin return min element in stack
func (this *MinStack) GetMin() int {
	return this.minStack[this.top-1]
}
