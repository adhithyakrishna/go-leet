package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result := make([]int, 0)

	calculateMax(root, result)
	return result[0]
}

func calculateMax(root *TreeNode, result []int) int {
	leftSubTree := max(0, calculateMax(root.Left, result))
	rightSubTree := max(0, calculateMax(root.Right, result))
	result[0] = max(result[0], leftSubTree+rightSubTree+root.Val)
	return max(leftSubTree, rightSubTree) + root.Val
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}
