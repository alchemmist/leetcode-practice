//go:build ignore

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
    mem := map[*ListNode]struct{}{}
    for head != nil {
        if _, ok := mem[head]; ok {
            return head
        }
        mem[head] = struct{}{}
        head = head.Next
    }
    return nil
}


func main() {
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: -4}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2

	entry := detectCycle(n1)
	fmt.Println(entry.Val)

	a := &ListNode{Val: 10}
	b := &ListNode{Val: 20}
	c := &ListNode{Val: 30}

	a.Next = b
	b.Next = c
	c.Next = nil

	fmt.Println(detectCycle(a))
}
