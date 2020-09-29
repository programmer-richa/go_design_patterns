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

	// Using Functional Builder
	pb := PersonBuilder{}
	r := pb.Name("Richa").Position("Software Developer").Build()
	fmt.Println(r)

	// Using Builder Facets
	cb := NewCustomerBuilder()
	cb.Works().Address("ABC Street").City("Delhi").Pincode("111111")
	cb.Lives().Address("XYZ Street").City("Delhi").Pincode("111111")
	c := cb.Build()
	fmt.Println(c)

	// b := &EmailBuilder{}
	// email := b.From("foo@bar.com").
	// 	To("bar@baz.com").
	// 	Subject("Meeting").
	// 	Body("Hello, do you want to meet?").
	// 	Build()
	// Builder Parameters
	SendEmail(func(b *EmailBuilder) {
		b.From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	})
}
