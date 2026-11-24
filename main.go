package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// MaxWidth returns the maximum width over all levels, where width counts
// the positions between a level's leftmost and rightmost nodes inclusive —
// including the null gaps between them. BFS carries heap-style position
// indexes, re-based per level so they never overflow on deep trees.
func MaxWidth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	type item struct {
		n   *TreeNode
		idx int
	}
	best := 0
	queue := []item{{root, 0}}
	for len(queue) > 0 {
		levelSize := len(queue)
		base := queue[0].idx // re-base: keeps indexes small
		if w := queue[levelSize-1].idx - base + 1; w > best {
			best = w
		}
		for i := 0; i < levelSize; i++ {
			it := queue[0]
			queue = queue[1:]
			idx := it.idx - base
			if it.n.Left != nil {
				queue = append(queue, item{it.n.Left, 2 * idx})
			}
			if it.n.Right != nil {
				queue = append(queue, item{it.n.Right, 2*idx + 1})
			}
		}
	}
	return best
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(1, tn(3, tn(5, nil, nil), tn(3, nil, nil)), tn(2, nil, tn(9, nil, nil)))
	fmt.Println(MaxWidth(root))
}
