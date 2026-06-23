package main

import (
	"fmt"
	"strings"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

// 思路：递归，判断左右子树是否对称
// 判断左右子树是否对称：如果左右子树都为空，则对称；如果左右子树有一个为空，则不对称；如果左右子树的值不相等，则不对称；如果左右子树的值相等，则继续判断左右子树的左右子树是否对称
func isSymmetric(root *TreeNode) bool {
	return isMirror(root.Left, root.Right)
}

func isMirror(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}

func main() {
	root := buildTreeFromLevel(parseTreeLevel("1,2,2,3,4,4,3"))
	fmt.Printf("PASS | symmetric=%v\n", isSymmetric(root))
}
