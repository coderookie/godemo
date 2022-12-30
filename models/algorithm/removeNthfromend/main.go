package main

import (
	"fmt"
)

type ListNode struct {
	Val		int
	Next	*ListNode
}

func createList(nums []int) *ListNode {
	head := new(ListNode)
	prev := head
	for _, v := range nums {
		node := &ListNode{
			Val: v,
			Next: nil,
		}
		prev.Next = node
		prev = node
	}

	return head.Next
}

func printList(head *ListNode) []int {
	ret := make([]int, 0)
	prev := head
	for prev != nil {
		ret = append(ret, prev.Val)
		prev = prev.Next
	}

	return ret
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n == 0 {
		return head
	}

	// 倒数, 换算成正数
	p := head
	l := 0
	for p != nil {
		l++
		p = p.Next
	}
	m := l - n + 1

	h := new(ListNode)
	h.Next = head
	prev := h
	i := 1
	for prev.Next != nil {
		if i == m {
			if prev.Next.Next != nil {
				prev.Next = prev.Next.Next
			} else {
				prev.Next = nil
			}
			break
		}
		i++
		prev = prev.Next
	}

	return h.Next
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	head := createList(nums)
	n := 2
	ret := removeNthFromEnd(head, n)
	fmt.Println(printList(ret))
}
