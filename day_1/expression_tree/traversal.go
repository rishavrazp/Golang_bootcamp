package expression_tree

import (
	"fmt"
	"github.com/Projects/Golang_bootcamp/day_1/model"
)

// Preorder traversal: root, left, right
func PreorderTraversal(root *model.Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Value, " ")
	PreorderTraversal(root.LeftNode)
	PreorderTraversal(root.RightNode)
}

// Postorder traversal: left, right, root
func PostorderTraversal(root *model.Node) {
	if root == nil {
		return
	}
	PostorderTraversal(root.LeftNode)
	PostorderTraversal(root.RightNode)
	fmt.Print(root.Value, " ")
}
