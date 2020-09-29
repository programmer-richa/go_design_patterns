package main

/***********************************************************************
* Using Builder Facets to i itialise complex element by using multiple builder classes
 */
type Customer struct {
	id, name                                 string
	homeAddress, homeCity, homePincode       string
	officeAddress, officeCity, officePincode string
}

// CustomerBuilder to chain initiatialisation of Customer struct
type CustomerBuilder struct {
	customer *Customer
}

func NewCustomerBuilder() *CustomerBuilder {
	return &CustomerBuilder{&Customer{}}
}

func (it *CustomerBuilder) Build() *Customer {
	return it.customer
}

// Works returns a pointer to CustomerJobBuilder
func (it *CustomerBuilder) Works() *CustomerJobBuilder {
	return &CustomerJobBuilder{*it}
}

// Lives returns a pointer to CustomerAddressBuilder
func (it *CustomerBuilder) Lives() *CustomerAddressBuilder {
	return &CustomerAddressBuilder{*it}
}

// CustomerJobBuilder initialise the parameters for customer job details
type CustomerJobBuilder struct {
	CustomerBuilder
}

func (it *CustomerJobBuilder) Address(address string) *CustomerJobBuilder {
	it.customer.officeAddress = address
	return it
}

func (it *CustomerJobBuilder) City(city string) *CustomerJobBuilder {
	it.customer.officeCity = city
	return it
}

func (it *CustomerJobBuilder) Pincode(pincode string) *CustomerJobBuilder {
	it.customer.officePincode = pincode
	return it
}

// CustomerAddressBuilder initialise the parameters for customer home details
type CustomerAddressBuilder struct {
	CustomerBuilder
}

func (it *CustomerAddressBuilder) Address(address string) *CustomerAddressBuilder {
	it.customer.homeAddress = address
	return it
}

func (it *CustomerAddressBuilder) City(city string) *CustomerAddressBuilder {
	it.customer.homeCity = city
	return it
}

func (it *CustomerAddressBuilder) Pincode(pincode string) *CustomerAddressBuilder {
	it.customer.homePincode = pincode
	return it
}
