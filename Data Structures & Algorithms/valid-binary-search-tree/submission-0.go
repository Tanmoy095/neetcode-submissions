/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {

 return isValid(root,math.MinInt64,math.MaxInt64)

    
}
func isValid(node *TreeNode,min int,max int) bool{
	if node==nil{
		return true
	}
	if node.Val<=min || node.Val>=max {
		return false

	}
	return isValid(node.Left,min,node.Val) && isValid(node.Right,node.Val,max)
}
