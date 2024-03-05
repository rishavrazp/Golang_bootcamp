package model

type EmployeeType int

const (
	FullTime EmployeeType = iota
	Contractor
	Freelancer
)

type Employee struct {
	Type EmployeeType
}

func NewEmployee(employeeType EmployeeType) *Employee {
	return &Employee{Type: employeeType}
}

func (e *Employee) CalculateSalary() float64 {
	switch e.Type {
	case FullTime:
		return 15000
	case Contractor:
		return 3000
	case Freelancer:
		return 2000
	default:
		return 0
	}
}
