package matrix

import (
	"github.com/Projects/Golang_bootcamp/day_1/model"
)

type matrixService struct{}

func NewMatrixService() IMatrixService {
	return &matrixService{}
}

func (s *matrixService) AddMatrices(matrix1, matrix2 *model.Matrix) (*model.Matrix, error) {
	return matrix1.Add(matrix2)
}
