package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func createList(data []int) *ListNode {
	h := new(ListNode)
	prev := h
	for _, v := range data {
		node := &ListNode{
			Val: v,
			Next: nil,
		}
		prev.Next = node
		prev = node
	}
	return h.Next
}

func printList(head *ListNode) []int {
	ret := make([]int, 0)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}

	return ret
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return new(ListNode)
	}
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var str1, str2, str3, str4 string

	for l1 != nil {
		str1 += strconv.Itoa(l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		str2 += strconv.Itoa(l2.Val)
		l2 = l2.Next
	}

	for i := len(str1) - 1; i >= 0; i-- {
		str3 += string(str1[i])
	}
	for i := len(str2) - 1; i >= 0; i-- {
		str4 += string(str2[i])
	}

	h := new(ListNode)
	prev := h
	i1, _ := strconv.Atoi(str3)
	i2, _ := strconv.Atoi(str4)
	i3 := i1 + i2
	str5 := strconv.Itoa(i3)

	var str string
	for i := len(str5) - 1; i >= 0; i-- {
		str += string(str5[i])
	}
	for _, v := range str {
		v1, _ := strconv.Atoi(string(v))
		node := &ListNode{
			Val: v1,
			Next: nil,
		}
		prev.Next = node
		prev = node
	}

	return h.Next
}

func main() {
	l1Int := []int{2, 4, 3}
	l2Int := []int{5, 6, 4}

	l1 := createList(l1Int)
	l2 := createList(l2Int)
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(printList(l3))
}