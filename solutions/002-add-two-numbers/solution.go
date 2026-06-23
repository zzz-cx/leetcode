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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cnt := 0
	dummy := new(ListNode)
	tail := dummy
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + cnt
		cnt = sum / 10
		sum = sum % 10
		tail.Next = &ListNode{Val: sum}
		tail = tail.Next
	}
	if cnt > 0 {
		tail.Next = &ListNode{Val: cnt}
	}
	return dummy.Next
}

func main() {
	l1 := sliceToList([]int{2, 4, 3})
	l2 := sliceToList([]int{5, 6, 4})
	got := listToSlice(addTwoNumbers(l1, l2))
	want := []int{7, 0, 8}
	status := "PASS"
	for i := range want {
		if got[i] != want[i] {
			status = "FAIL"
		}
	}
	fmt.Printf("%s | result=%v (expected %v)\n", status, got, want)
}
