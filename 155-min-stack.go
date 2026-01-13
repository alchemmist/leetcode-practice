//go:build ignore

package main

type Node struct {
	value int
	next  *Node
}

type MinStack struct {
	top    *Node
	length int
	mins   []*Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if this.top == nil {
		this.top = &Node{value: val}
		this.mins = append(this.mins, this.top)
	} else {
		this.top = &Node{value: val, next: this.top}
		this.length++
		if this.top.value <= this.mins[len(this.mins)-1].value {
			this.mins = append(this.mins, this.top)
		}
	}
}

func (this *MinStack) Pop() {
	if this.mins[len(this.mins)-1] == this.top {
		this.mins = this.mins[:len(this.mins)-1]
	}
	this.top = this.top.next
}

func (this *MinStack) Top() int {
	return this.top.value
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1].value
}
