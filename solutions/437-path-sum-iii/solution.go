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

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	Val := root.Val
	if Val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-Val)
	res += rootSum(root.Right, targetSum-Val)
	return
}
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func main() {
	root := buildTreeFromLevel(parseTreeLevel("1,2,3,4,5,6,7"))
	fmt.Printf("PASS | paths=%d\n", pathSum(root, 3))
}
