package main

/**************************************************************************
* Open closed Principle implementation
	In the following example, Specification interface lists a method : IsSatisfied,
	this method can be defined by different structs to define the criteria of filteration.
	Hence, it supports extention. However, there is a rare need for modification of this interface.

Similarly, The Filter struct has a method FilterProduct, that filter the product that matches a criteria.
FilterBook and other filters can be created just by copying similar functionality.

*/

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.Color() == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.Size() == spec.size
}

type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) &&
		spec.second.IsSatisfied(p)
}

type Filter struct{}

func (f *Filter) FilterProduct(products []*Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(v) {
			result = append(result, products[i])
		}
	}
	return result
}
