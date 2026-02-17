/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */




func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}

func dfs(node *TreeNode, maxSoFar int) {
    if maxSoFar <= node.Val {
        numGoodNodes += 1
    }

    if node.Right != nil {
        dfs(node.Right, max(maxSoFar, node.Val))
    }

    if node.Left != nil {
        dfs(node.Left, max(maxSoFar, node.Val))
    }
}

var numGoodNodes int
func goodNodes(root *TreeNode) int {
    numGoodNodes = 0
    dfs(root, math.MinInt)
    return numGoodNodes
}
