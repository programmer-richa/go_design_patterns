package main

import "fmt"

func main() {
	// Builder Design pattern Implementation
	emp := NewEmployee().
		SetID("101").
		SetName("Raman").
		SetBranch("Head Office").
		SetSalary(10000)
	fmt.Println(emp)
}
