package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func createList(a1 []int) *ListNode {
	head := new(ListNode)
	p := head
	for _, v := range a1 {
		node := &ListNode{
			Val: v,
			Next: nil,
		}
		p.Next = node
		p = node
	}
	return head.Next
}

func printList(l *ListNode) []int {
	ret := make([]int, 0)
	for l != nil {
		ret = append(ret, l.Val)
		l = l.Next
	}
	return ret
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1 == nil && list2 == nil {
		return nil
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func main() {
	a1 := []int{1, 2, 4}
	a2 := []int{1, 3, 4}
	l1 := createList(a1)
	l2 := createList(a2)
	ret := mergeTwoLists(l1, l2)
	fmt.Println(printList(ret))
}
