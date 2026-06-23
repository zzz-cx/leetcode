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
	var prev *ListNode
	for head != nil {
		nxt := head.Next
		head.Next = prev
		prev = head
		head = nxt
	}
	return prev
}

// isPalindrome 简单版（推荐面试先写）：链表值拷到切片，左右指针比较
// 时间 O(n)，空间 O(n)。代码短、不易错。
func isPalindrome(head *ListNode) bool {
	vals := make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		if vals[i] != vals[j] {
			return false
		}
	}
	return true
}

// isPalindromeO1 进阶版：快慢指针找中点 + 反转后半 + 比较，空间 O(1)
func isPalindromeO1(head *ListNode) bool {
	if head == nil {
		return true
	}
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)
	p1, p2 := head, secondHalfStart
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1, p2 = p1.Next, p2.Next
	}
	firstHalfEnd.Next = reverseList(secondHalfStart) // 恢复链表（面试可省略）
	return true
}

func endOfFirstHalf(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func main() {
	head := sliceToList([]int{1, 2, 2, 1})
	fmt.Printf("PASS | %v\n", isPalindrome(head))
}
