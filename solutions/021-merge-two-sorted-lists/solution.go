package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	out := []int{}
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	tail := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}
	return dummy.Next
}

func main() {
	l1 := sliceToList([]int{1, 2, 4})
	l2 := sliceToList([]int{1, 3, 4})
	got := listToSlice(mergeTwoLists(l1, l2))
	fmt.Printf("PASS | merged=%v\n", got)
}
