package main

/******************************************************************************
*	Functional Builder implementation
 */
type Person struct {
	name, position string
}

// personModifier has a type of function that accepts a *Person
// functional builder
type personModifier func(*Person)

// PersonBuilder holds a list of functions that accepts a *Person
type PersonBuilder struct {
	actions []personModifier
}

func (b *PersonBuilder) Name(name string) *PersonBuilder {
	// append an anonymous function that specifies name of a person
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

// extend PersonBuilder
func (b *PersonBuilder) Position(position string) *PersonBuilder {
	// append an anonymous function that specifies position of a person
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}
func (b *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}
	return &p
}
