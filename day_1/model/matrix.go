package model

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Rows     int
	Cols     int
	Elements [][]int
}

func NewMatrix(rows, cols int) *Matrix {
	elements := make([][]int, rows)
	for i := range elements {
		elements[i] = make([]int, cols)
	}
	return &Matrix{
		Rows:     rows,
		Cols:     cols,
		Elements: elements,
	}
}

func (m *Matrix) GetRows() int {
	return m.Rows
}

func (m *Matrix) GetCols() int {
	return m.Cols
}

func (m *Matrix) SetElement(i, j, value int) error {
	if i < 0 || i >= m.Rows || j < 0 || j >= m.Cols {
		return fmt.Errorf("index out of range")
	}
	m.Elements[i][j] = value
	return nil
}

func (m *Matrix) Add(other *Matrix) (*Matrix, error) {
	if m.Rows != other.Rows || m.Cols != other.Cols {
		return nil, fmt.Errorf("matrix dimensions mismatch")
	}
	result := NewMatrix(m.Rows, m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Elements[i][j] = m.Elements[i][j] + other.Elements[i][j]
		}
	}
	return result, nil
}

func (m *Matrix) PrintAsJSON() {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error marshaling matrix to JSON:", err)
		return
	}
	fmt.Println("\n", string(jsonBytes))
}
