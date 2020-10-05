package main

type Staff struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

// functional
func NewStaff(role int) *Staff {
	switch role {
	case Developer:
		return &Staff{"", "Developer", 60000}
	case Manager:
		return &Staff{"", "Manager", 80000}
	default:
		panic("unsupported role")
	}
}
