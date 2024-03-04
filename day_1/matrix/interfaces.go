package matrix

import "github.com/Projects/Golang_bootcamp/day_1/model"

type IMatrixService interface {
	AddMatrices(matrix1, matrix2 *model.Matrix) (*model.Matrix, error)
}
