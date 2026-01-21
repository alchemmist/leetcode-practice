//go:build ignore

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
    fast := head
    slow := head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if slow == fast {
            slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
        }
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
