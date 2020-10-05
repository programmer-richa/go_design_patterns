package main

type Employee struct {
	name, position string
	annualIncome   int
}

func (e *Employee) Name() string {
	return e.name
}

func (e *Employee) Position() string {
	return e.position
}

func (e *Employee) AnnualIncome() int {
	return e.annualIncome
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) SetPosition(position string) {
	e.position = position
}

func (e *Employee) SetAnnualIncome(annualIncome int) {
	e.annualIncome = annualIncome
}

// functional approach that returns a function to initialise variables based on roles.
// It is idomatic way
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// The below approach is not idomatic
// structural approach
type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}
