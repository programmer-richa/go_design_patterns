package main

import "fmt"

func main() {
	// calling factory function to retrieve pointer to person
	p := NewPerson("Rahul", 23)
	fmt.Println(p)
	// Returning Interface instead of pointer to teacher struct
	t := NewTeacher("Aman", "MCA")
	// calling SetName method to modify data
	// as the properties of teacher struct can not be accessed using reference of interface
	t.SetName("Preeti")
	fmt.Println(t.Name())

	developerFactory := NewEmployeeFactory("Developer", 60000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	developer := developerFactory("Adam")
	developer.SetName("Raman")
	fmt.Println(developer)

	manager := managerFactory("Jane")
	fmt.Println(manager)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	// can modify post-hoc
	bossFactory.AnnualIncome = 110000
	boss := bossFactory.Create("Sam")
	fmt.Println(boss)
	// Prototype Factory
	m := NewStaff(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
