package main

type Person struct {
	name        string
	age         int
	nationality string
}

// Factory function for Person
func NewPerson(name string, age int) *Person {
	return &Person{name, age, "Indian"}
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) SetAge(age int) {
	p.age = age
}
func (p *Person) SetNationality(nationality string) {
	p.nationality = nationality
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) Age() int {
	return p.age
}
func (p *Person) Nationality() string {
	return p.nationality
}
