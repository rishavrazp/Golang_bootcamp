package main

import (
	"fmt"
	"github.com/Projects/Golang_bootcamp/day_1/controller"
	"github.com/Projects/Golang_bootcamp/day_1/matrix"
	"github.com/Projects/Golang_bootcamp/day_1/model"
)

// setErrorHandler is a helper function to handle errors during setting an element
func setErrorHandler(err error) {
	if err != nil {
		fmt.Println("Error setting element:", err)
		panic(err)
	}
}

func main() {
	matrix1 := model.NewMatrix(2, 2)
	setErrorHandler(matrix1.SetElement(0, 0, 1))
	setErrorHandler(matrix1.SetElement(0, 1, 2))
	setErrorHandler(matrix1.SetElement(1, 0, 3))
	setErrorHandler(matrix1.SetElement(1, 1, 4))

	matrix2 := model.NewMatrix(2, 2)
	setErrorHandler(matrix2.SetElement(0, 0, 5))
	setErrorHandler(matrix2.SetElement(0, 1, 6))
	setErrorHandler(matrix2.SetElement(1, 0, 7))
	setErrorHandler(matrix2.SetElement(1, 1, 8))

	// Create Matrix Service
	service := matrix.NewMatrixService()

	// Create Matrix Controller with Matrix Service
	matrixController := controller.NewMatrixController(service)

	// Call Controller Method
	result, err := matrixController.AddMatrices(matrix1, matrix2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	result.PrintAsJSON()
}
