package main

import "fmt"

/*************************************************************************
The principle defines that objects of a superclass shall be replaceable with objects
 of its subclasses without breaking the application. That requires the objects of your
 subclasses to behave in the same way as the objects of your superclass.
*/

// Sizeable interface outlines the method for the elements that have height and width.
// The core idea behind this interface is that the height and width should be assigned a value separately.
type Sizeable interface {
	SetHeight(height int)
	SetWidth(width int)
	Height() int
	Width() int
}

// Rectangle struct implements Sizeable interface
type Rectangle struct {
	height int
	width  int
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) Height() int {
	return r.height
}

func (r *Rectangle) Width() int {
	return r.width
}

// Square struct contains Rectangle struct variable.
// The core idea behind the Sizeable Interface remains unchanged with this arrangement.
// It is because for using square as an rectangle, the height and width values can be modified separately.
type Square struct {
	Rectangle
}

func NewSquare(side int) *Square {
	rectange := Rectangle{}
	rectange.SetHeight(side)
	rectange.SetWidth(side)
	return &Square{rectange}
}

func LiskovSubstitutionPrinciple() {
	s := NewSquare(5)
	// Modifies height value only.
	// The value of width remains unchanged.
	s.SetHeight(10)
	fmt.Println(s)
}
