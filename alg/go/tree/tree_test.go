package tree

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)

	root.Left.Left = NewNode(4)
	root.Right.Left = NewNode(5)

	fn := func(root *TreeNode) {
		fmt.Print(root.Val, " ")
	}

	Preorder(root, fn)

	t.Logf("\n")

	Inorder(root, fn)
	t.Logf("\n")

	Postorder(root, fn)
	t.Logf("\n")

	BreadthTraverse(root, fn)
	t.Logf("\n")

	DepthTraverse(root, fn)
}

func TestDepth(t *testing.T) {

}
