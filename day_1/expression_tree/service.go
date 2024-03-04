package expression_tree

import "github.com/Projects/Golang_bootcamp/day_1/model"

type ExpressionTree struct{}

func (e *ExpressionTree) PreorderTraversal(root *model.Node) {
	PreorderTraversal(root)
}

func (e *ExpressionTree) PostorderTraversal(root *model.Node) {
	PostorderTraversal(root)
}
