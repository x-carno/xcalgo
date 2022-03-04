package trees

import (
	"math"

	"github.com/x-carno/xcalgo/listnode"
)

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

// given binary tree
// 3
// / \
// 9  20
//  /  \
// 15   7
// return
// [
//   [3],
//   [9,20],
//   [15,7]
// ]
func (root *TreeNode) LevelOrder2D() [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	queue := make([][]*TreeNode, 0)
	queue = append(queue, []*TreeNode{root})
	for len(queue) != 0 {
		layer := queue[0]
		layerArr := make([]int, 0)
		nextLayer := make([]*TreeNode, 0)
		for _, node := range layer {
			layerArr = append(layerArr, node.Val)
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}
		ret = append(ret, layerArr)
		if len(nextLayer) != 0 {
			queue = append(queue, nextLayer)
		}
		queue = queue[1:]
	}

	return ret
}

func (root *TreeNode) Height() int {
	height := 0
	if root == nil {
		return height
	}
	queue := make([][]*TreeNode, 0)
	queue = append(queue, []*TreeNode{root})
	for len(queue) != 0 {
		layer := queue[0]
		nextLayer := make([]*TreeNode, 0)
		for _, node := range layer {
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}
		if len(nextLayer) != 0 {
			queue = append(queue, nextLayer)
		}
		queue = queue[1:]
		height++
	}
	return height
}

func (root *TreeNode) IsBalanced() bool {
	if root == nil {
		return true
	}

	hleft := root.Left.Height()
	hright := root.Right.Height()

	return math.Abs(float64(hleft-hright)) <= 1 && root.Left.IsBalanced() && root.Right.IsBalanced()
}

// given binary tree
// 3
// / \
// 9  20
//  /  \
// 15   7
// return
// [
//   [3],
//   [20,9],
//   [15,7]
// ]
func (root *TreeNode) LevelOrder2DReverseOdd() [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	queue := make([][]*TreeNode, 0)
	queue = append(queue, []*TreeNode{root})
	layerNum := 0
	for len(queue) != 0 {
		layer := queue[0]
		layerArr := make([]int, 0)
		nextLayer := make([]*TreeNode, 0)
		if layerNum%2 == 0 {
			for i := 0; i < len(layer); i++ {
				layerArr = append(layerArr, layer[i].Val)
			}
		} else {
			for i := len(layer) - 1; i >= 0; i-- {
				layerArr = append(layerArr, layer[i].Val)
			}
		}
		layerNum++
		for _, node := range layer {
			// layerArr = append(layerArr, node.Val)
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}
		ret = append(ret, layerArr)
		if len(nextLayer) != 0 {
			queue = append(queue, nextLayer)
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

//input：
//      4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9
// output：
//      4
//    /   \
//   7     2
//  / \   / \
// 9   6 3   1
func (root *TreeNode) MirrorTree() *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right.MirrorTree(), root.Left.MirrorTree()
	return root
}

// 对称二叉树
func (root *TreeNode) IsSymmetric() bool {
	if root == nil {
		return true
	}
	return isSymmetric(root.Left, root.Right)
}

func isSymmetric(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
		return false
	}
	return n1.Val == n2.Val && isSymmetric(n1.Left, n2.Right) && isSymmetric(n1.Right, n2.Left)
}

// whether b is a subtree of a
func (A *TreeNode) IsSubStructure(B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return A.hasPrefixSequence(B) || A.Left.IsSubStructure(B) || A.Right.IsSubStructure(B)
}

// determine whether B is a prefix sequence of A
// nil is a sub sequence of any other tree
// given A
//      4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9
// given B
//      4
//    /
//   2
// return true
func (A *TreeNode) hasPrefixSequence(B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val != B.Val {
		return false
	}
	return A.Left.hasPrefixSequence(B.Left) && A.Right.hasPrefixSequence(B.Right)
}

func (root *TreeNode) IsSubPath(head *listnode.ListNode) bool {
	if root == nil || head == nil {
		return false
	}
	return root.hasPrefixPath(head) || root.Left.IsSubPath(head) || root.Right.IsSubPath(head)
}

func (root *TreeNode) hasPrefixPath(head *listnode.ListNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if root.Val != head.Val {
		return false
	}
	return root.Left.hasPrefixPath(head.Next) || root.Right.hasPrefixPath(head.Next)
}
