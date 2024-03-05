package main

import (
	"fmt"
	"github.com/Projects/Golang_bootcamp/day_1/controller"
	"github.com/Projects/Golang_bootcamp/day_1/model"
)

func EmployeeMain() {

	//Full-time Employee
	fullTimeEmployee := model.NewEmployee(model.FullTime)
	fullTimeController := controller.NewEmployeeController(fullTimeEmployee)
	fmt.Println("\nFull-time Employee Salary:", fullTimeController.CalculateSalary())

	//Contractor
	contractor := model.NewEmployee(model.Contractor)
	contractorController := controller.NewEmployeeController(contractor)
	fmt.Println("Contractor Salary:", contractorController.CalculateSalary())

	//Freelancer
	freelancer := model.NewEmployee(model.Freelancer)
	freelancerController := controller.NewEmployeeController(freelancer)
	fmt.Println("Freelancer Salary:", freelancerController.CalculateSalary())
}
