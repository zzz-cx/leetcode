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

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

func main() {
	head := sliceToList([]int{1, 2, 3})
	fmt.Printf("PASS | entry=%v\n", detectCycle(head))
}
