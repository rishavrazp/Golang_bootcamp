package controller

import (
	"github.com/Projects/Golang_bootcamp/day_1/employee"
)

type EmployeeController struct {
	service employee.IEmployeeService
}

func NewEmployeeController(service employee.IEmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (c *EmployeeController) CalculateSalary() float64 {
	return c.service.CalculateSalary()
}
