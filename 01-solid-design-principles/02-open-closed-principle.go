package main

import "fmt"

/************************************************
* Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification
 */

func openClosedPrinciple() {
	apple := NewProduct().
		SetTitle("Apple").
		SetColor(Green).
		SetSize(Small)

	tree := NewProduct().
		SetTitle("Tree").
		SetColor(Green).
		SetSize(Large)

	house := NewProduct().
		SetTitle("House").
		SetColor(Brown).
		SetSize(Large)
	products := []*Product{apple, tree, house}
	fmt.Print("Green products (new):\n")
	greenSpec := ColorSpecification{Green}
	f := Filter{}
	for _, v := range f.FilterProduct(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.Title())
	}

	largeSpec := SizeSpecification{Large}
	brownSpec := ColorSpecification{Brown}
	largeBrownSpec := AndSpecification{largeSpec, brownSpec}
	fmt.Print("Large brown items:\n")
	for _, v := range f.FilterProduct(products, largeBrownSpec) {
		fmt.Printf(" - %s is large and brown\n", v.Title())
	}
}
