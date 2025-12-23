//go:build ignore

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var mergedList *ListNode = nil
	var last *ListNode = nil
	var state string
	for {
		var l *ListNode
		if list1.Val <= list2.Val {
			l = list1
			list1 = list1.Next
			state = "1"
		} else {
			l = list2
			list2 = list2.Next
			state = "2"
		}

		if mergedList == nil {
			mergedList = l
		}

		if last == nil {
			last = l
			if l.Next == nil {
				break
			}
			continue
		}

		last.Next = l
		last = l

		if l.Next == nil {
			break
		}
	}
	if state == "1" {
		last.Next = list2
	} else {
		last.Next = list1
	}
	return mergedList
}

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	cur := head
	for i := 1; i < len(vals); i++ {
		cur.Next = &ListNode{Val: vals[i]}
		cur = cur.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	tests := []struct {
		a []int
		b []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}},
		{[]int{}, []int{}},
		{[]int{}, []int{0, 1}},
		{[]int{5}, []int{1, 2, 3}},
		{[]int{2}, []int{1}},
	}

	for i, t := range tests {
		fmt.Printf("Test %d:\n", i+1)

		l1 := buildList(t.a)
		l2 := buildList(t.b)

		fmt.Print("list1: ")
		printList(l1)
		fmt.Print("list2: ")
		printList(l2)

		res := mergeTwoLists(l1, l2)

		fmt.Print("merged: ")
		printList(res)
		fmt.Println()
	}
}
