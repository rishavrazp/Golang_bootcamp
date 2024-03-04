package controller

import (
	"github.com/Projects/Golang_bootcamp/day_1/matrix"
	"github.com/Projects/Golang_bootcamp/day_1/model"
)

type MatrixController struct {
	service matrix.IMatrixService
}

func NewMatrixController(service matrix.IMatrixService) *MatrixController {
	return &MatrixController{service: service}
}

func (c *MatrixController) AddMatrices(matrix1, matrix2 *model.Matrix) (*model.Matrix, error) {
	return c.service.AddMatrices(matrix1, matrix2)
}
