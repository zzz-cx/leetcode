package main

import (
	"fmt"
	"strings"
	"strconv"
)

func buildTreeFromLevel(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]
		if i < len(values) && values[i] != nil {
			node.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func parseTreeLevel(s string) []interface{} {
	s = strings.Trim(s, "[]")
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]interface{}, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "null" || p == "" {
			out = append(out, nil)
		} else {
			v, _ := strconv.Atoi(p)
			out = append(out, v)
		}
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路：递归，先遍历左子树，再遍历根节点，再遍历右子树
//这里使用外部out＋内部递归函数的形式
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	out := inorderTraversal(root.Left)
	out = append(out, root.Val)
	right := inorderTraversal(root.Right)
	out = append(out, right...)
	return out
}

func main() {
	root := buildTreeFromLevel(parseTreeLevel("1,2,3,4,5,6,7"))
	got := inorderTraversal(root)
	want := []int{4, 2, 5, 1, 6, 3, 7}
	status := "PASS"
	for i := range want {
		if got[i] != want[i] {
			status = "FAIL"
		}
	}
	fmt.Printf("%s | inorder=%v\n", status, got)
}
