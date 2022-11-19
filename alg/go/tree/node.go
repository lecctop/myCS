package tree

import (
	"github.lecctop.myCS/alg/internal/utils"
)

const (
	preorderType = 1 + iota
	inorderType
	postorderType
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func Preorder(root *TreeNode, f func(root *TreeNode)) {
	traverse(root, preorderType, f)
}

func Inorder(root *TreeNode, f func(root *TreeNode)) {
	traverse(root, inorderType, f)
}

func Postorder(root *TreeNode, f func(root *TreeNode)) {
	traverse(root, postorderType, f)
}

func traverse(root *TreeNode, orderType int, f func(root *TreeNode)) {
	if root == nil {
		return
	}
	if orderType == preorderType {
		f(root)
	}
	traverse(root.Left, orderType, f)
	if orderType == inorderType {
		f(root)
	}
	traverse(root.Right, orderType, f)
	if orderType == postorderType {
		f(root)
	}

}

func BreadthTraverse(root *TreeNode, f func(root *TreeNode)) {
	if root == nil {
		return
	}

	q := make([]*TreeNode, 0, 8)
	q = append(q, root)
	for len(q) > 0 {
		n := q[0]

		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
		q = q[1:]

		f(n)

	}
}

func DepthTraverse(root *TreeNode, f func(root *TreeNode)) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0, 8)
	stack = append(stack, root)
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
		if n.Left != nil {
			stack = append(stack, n.Left)
		}

		f(n)
	}
}

// MaxDepth 最大深度
//
// 104. Maximum Depth of Binary Tree
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
//
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	return 1 + utils.MaxInt(left, right)
}

// IsBalanced 高度是否平衡
// 左右子树高度不大于1
//
// 110. Balanced Binary Tree
// https://leetcode.com/problems/balanced-binary-tree/
//
func IsBalanced(root *TreeNode) bool {
	_, _, rs := balanced(root)
	return rs
}

func balanced(root *TreeNode) (left int, right int, balance bool) {
	if root == nil {
		return 0, 0, true
	}

	l1, r1, b1 := balanced(root.Left)
	l2, r2, b2 := balanced(root.Right)

	left = utils.MaxInt(l1, r1)
	right = utils.MaxInt(l2, r2)

	if root.Left != nil {
		left++
	}
	if root.Right != nil {
		right++
	}
	balance = b1 && b2 && utils.AbsInt(left-right) <= 1

	return left, right, balance
}

// IsSameTree 是否相同的树
//
// 100. Same Tree
// https://leetcode.com/problems/same-tree/
//
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return same(p, q)
}

func same(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	left := same(p.Left, q.Left)
	right := same(p.Right, q.Right)

	return left && right
}

// IsSymmetric 对称二叉树
//
// 101. Symmetric Tree
// https://leetcode.com/problems/symmetric-tree/
//
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return symmetric(root.Left, root.Right)
}

func symmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val == right.Val {
		is1 := symmetric(left.Left, right.Right)
		is2 := symmetric(left.Right, right.Left)

		return is1 && is2
	}

	return false
}

// SumOfLeftLeaves 左叶子结点的和
//
// 404. Sum of Left Leaves
// https://leetcode.com/problems/sum-of-left-leaves/
//
func SumOfLeftLeaves(root *TreeNode) int {
	return sumLeaves(root.Left, true) + sumLeaves(root.Right, false)
}

func sumLeaves(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil && isLeft {
		return root.Val
	}

	v1 := sumLeaves(root.Left, true)
	v2 := sumLeaves(root.Right, false)

	return v1 + v2
}
