//go:build ignore

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	curr := head

	for curr != nil {
		next := curr.Next
		pos := dummy

		for pos.Next != nil && pos.Next.Val < curr.Val {
			pos = pos.Next
		}

		curr.Next = pos.Next
		pos.Next = curr
		curr = next
	}

	return dummy.Next
}

func main() {
	l := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	l = insertionSortList(l)
	for l != nil {
		fmt.Printf("%d -> ", l.Val)
		l = l.Next
	}
}
