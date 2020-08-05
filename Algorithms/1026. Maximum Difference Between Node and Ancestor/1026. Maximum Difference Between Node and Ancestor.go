package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
	res := 0
	dfsAncestorDiff(root, &res)
	return res
}

func dfsAncestorDiff(root *TreeNode, res *int) (int, int) {
	if root == nil {
		return -1, -1
	}
	leftMax, leftMin := dfsAncestorDiff(root.Left, res)
	if leftMax == -1 && leftMin == -1 {
		leftMax = root.Val
		leftMin = root.Val
	}
	rightMax, rightMin := dfsAncestorDiff(root.Right, res)
	if rightMax == -1 && rightMin == -1 {
		rightMax = root.Val
		rightMin = root.Val
	}
	*res = max(*res, max(abs(root.Val-min(leftMin, rightMin)), abs(root.Val-max(leftMax, rightMax))))
	return max(leftMax, max(rightMax, root.Val)), min(leftMin, min(rightMin, root.Val))
}
