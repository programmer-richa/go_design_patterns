package main

type Employee struct {
	id     string
	name   string
	branch string
	salary int
}

func NewEmployee() *Employee {
	return &Employee{}
}

func (e *Employee) SetID(id string) *Employee {
	e.id = id
	return e
}

func (e *Employee) SetName(name string) *Employee {
	e.name = name
	return e
}

func (e *Employee) SetBranch(branch string) *Employee {
	e.branch = branch
	return e
}

func (e *Employee) SetSalary(salary int) *Employee {
	e.salary = salary
	return e
}

func (e *Employee) ID() string {
	return e.id
}

func (e *Employee) Name() string {
	return e.name
}

func (e *Employee) Branch() string {
	return e.branch
}

func (e *Employee) Salary() int {
	return e.salary
}
