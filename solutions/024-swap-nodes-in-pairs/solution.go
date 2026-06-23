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

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy
	current := head
	for current != nil && current.Next != nil {
		next := current.Next
		current.Next = next.Next
		next.Next = current
		prev.Next = next
		prev = current
		current = current.Next
	}
	return dummy.Next
}

func main() {
	head := sliceToList([]int{1, 2, 3, 4})
	got := listToSlice(swapPairs(head))
	fmt.Printf("PASS | %v\n", got)
}
