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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first := dummy
	second := dummy
	for i := 0; i < n+1; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next

}

func main() {
	head := sliceToList([]int{1, 2, 3, 4, 5})
	got := listToSlice(removeNthFromEnd(head, 2))
	want := []int{1, 2, 3, 5}
	status := "PASS"
	if len(got) != len(want) {
		status = "FAIL"
	}
	fmt.Printf("%s | %v (expected %v)\n", status, got, want)
}
