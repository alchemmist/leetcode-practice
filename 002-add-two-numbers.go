//go:build ignore

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func add(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	sum := carry
	var l1_ *ListNode = nil
	var l2_ *ListNode = nil
	if l1 != nil {
		sum += l1.Val
		l1_ = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		l2_ = l2.Next
	}

	return &ListNode{
		Val:  sum % 10,
		Next: add(l1_, l2_, sum/10),
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return add(l1, l2, 0)
}

func main() {
	l1 := &ListNode{3, nil}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	res := addTwoNumbers(l1, l2)

	for res != nil {
		fmt.Print(res.Val)
		if res.Next != nil {
			fmt.Print(" -> ")
		}
		res = res.Next
	}
	fmt.Println()
}
