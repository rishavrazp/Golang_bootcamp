package controller

import (
	"github.com/Projects/Golang_bootcamp/day_1/expression_tree"
	"github.com/Projects/Golang_bootcamp/day_1/model"
)

type ExpressionTreeController struct {
	service expression_tree.IExpressionTreeService
}

func NewExpressionTreeController(service expression_tree.IExpressionTreeService) *ExpressionTreeController {
	return &ExpressionTreeController{service: service}
}

func (c *ExpressionTreeController) PreorderTraversal(root *model.Node) {
	c.service.PreorderTraversal(root)
}

func (c *ExpressionTreeController) PostorderTraversal(root *model.Node) {
	c.service.PostorderTraversal(root)
}
