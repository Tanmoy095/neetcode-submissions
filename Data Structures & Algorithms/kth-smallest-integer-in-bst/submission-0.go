/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "slices"


func kthSmallest(root *TreeNode, k int) int {

	if root == nil {
		return 0
	}

	output := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {

		curr := queue[0]
		queue = queue[1:]
		if curr.Val != 0 {
			output = append(output, curr.Val)
		}

		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}

	}
	slices.Sort(output)
	target := output[k-1] // kth smallest element (1-indexed)
	return target

}
