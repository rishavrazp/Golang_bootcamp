package main

import (
	"github.com/Projects/Golang_bootcamp/day_1/controller"
	"github.com/Projects/Golang_bootcamp/day_1/expression_tree"
	"github.com/Projects/Golang_bootcamp/day_1/model"
)

func ExpressionTreeMain() {

	// Build the expression tree: a + b - c
	a := model.NewNode("a")
	b := model.NewNode("b")
	c := model.NewNode("c")

	minus := model.NewBinaryExpressionTree("-", b, c)
	plus := model.NewBinaryExpressionTree("+", a, minus)

	// Create Service for Expression Tree
	service := &expression_tree.ExpressionTree{}

	// Create Controller with service
	expressionTreeController := controller.NewExpressionTreeController(service)

	// Perform Preorder Traversal
	println("\nPreorder Traversal:")
	expressionTreeController.PreorderTraversal(plus)
	println()

	// Perform Postorder Traversal
	println("\nPostorder Traversal:")
	expressionTreeController.PostorderTraversal(plus)
	println()
}
