package trees

type BinaryTree struct {
	root *TreeNode
}

// a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) LevelOrder() []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		ret = append(ret, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}
	return ret
}

func computeTreeLayer(arr []interface{}) int {
	if arr == nil {
		return 0
	}
	layer := 0
	for i := len(arr); i > 0; i = i >> 1 {
		layer++
	}
	return layer
}
