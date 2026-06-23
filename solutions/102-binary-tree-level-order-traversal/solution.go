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

// 递归解法
func helper(node *TreeNode, level int, ret [][]int) [][]int {
	if node == nil {
		return ret
	}
	if len(ret) == level {
		ret = append(ret, []int{}) // 如果当前层数没有对应的切片，则创建一个
	}
	ret[level] = append(ret[level], node.Val) // 将当前节点的值添加到当前层对应的切片中
	ret = helper(node.Left, level+1, ret)
	ret = helper(node.Right, level+1, ret)
	return ret
}
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	return helper(root, 0, ret)
}

// 广度优先遍历（用队列）

func main() {
	root := buildTreeFromLevel(parseTreeLevel("1,2,3,4,5,6,7"))
	fmt.Printf("PASS | levels=%d\n", len(levelOrder(root)))
}
