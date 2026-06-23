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

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func main() {
	head := sliceToList([]int{1, 2, 3, 4, 5})
	got := listToSlice(reverseList(head))
	want := []int{5, 4, 3, 2, 1}
	status := "PASS"
	for i := range want {
		if got[i] != want[i] {
			status = "FAIL"
		}
	}
	fmt.Printf("%s | reversed=%v\n", status, got)
}
