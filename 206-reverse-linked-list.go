//go:build ignore

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	l := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return l
}

func main() {
	n3 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}

	head := reverseList(n1)

	for cur := head; cur != nil; cur = cur.Next {
		println(cur.Val)
	}
}
