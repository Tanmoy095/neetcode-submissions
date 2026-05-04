/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root==nil{
		return nil
	}
	queue := []*TreeNode{root}
	output := [][]int{}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelValues := []int{}

		for i := 0; i < levelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:]

			
				levelValues = append(levelValues, currentNode.Val)

				if currentNode.Left != nil {
					queue = append(queue, currentNode.Left)
				}
				if currentNode.Right != nil {
					queue = append(queue, currentNode.Right)
				}
			

		}
		output = append(output, levelValues)
	}
	return output

}