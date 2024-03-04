package expression_tree

import "github.com/Projects/Golang_bootcamp/day_1/model"

type IExpressionTreeService interface {
	PreorderTraversal(root *model.Node)
	PostorderTraversal(root *model.Node)
}
