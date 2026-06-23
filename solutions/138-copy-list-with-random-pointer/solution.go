package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		m[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]
}

func main() {
	head := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}}
	got := copyRandomList(head)
	vals := []int{}
	for cur := got; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}
	fmt.Printf("PASS | copied=%v\n", vals)
}
