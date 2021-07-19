/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}
	node := root
	for node != nil {
		if node.Val < p.Val && node.Val < q.Val {
			node = node.Right
		}
		if node.Val > p.Val && node.Val > q.Val {
			node = node.Left
		} else {
			return node
		}
	}
	return nil
}
