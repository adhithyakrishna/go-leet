package main

type QueueVal struct {
	data []*TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewQueueVal() QueueVal {
	q := QueueVal{}
	return q
}

func (q *QueueVal) insert(v *TreeNode) {
	q.data = append(q.data, v)
}

func (q *QueueVal) pop() *TreeNode {
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	dataQueueVal := NewQueueVal()

	// result := make([][]int, 0)

	var result [][]int
	var resArr []int

	if root == nil {
		return result
	}

	result = append(result, []int{root.Val})

	dataQueueVal.insert(root)

	for len(dataQueueVal.data) > 0 {
		datalen := len(dataQueueVal.data)
		resArr = nil
		for datalen > 0 {
			curr := dataQueueVal.pop()

			if curr.Left != nil {
				dataQueueVal.insert(curr.Left)
				resArr = append(resArr, curr.Left.Val)
			}

			if curr.Right != nil {
				dataQueueVal.insert(curr.Right)
				resArr = append(resArr, curr.Right.Val)
			}

			datalen--

		}
		if len(resArr) > 0 {
			result = append(result, resArr)
		}
	}
	return result
}
