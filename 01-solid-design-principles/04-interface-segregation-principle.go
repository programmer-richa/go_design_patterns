package main

import "fmt"

/********************************************************************************
In the field of software engineering, the interface-segregation principle (ISP)
states that no client should be forced to depend on methods it does not use

In the following example Rect struct need not to implement Volume interface as it
is not required in it.
*/

type Area interface {
	Area() int
}

type Volume interface {
	SetBreadth(breadth int)
	Breadth() int
	Volume() int
}

// Rect struct
type Rect struct {
	height int
	width  int
}

func (r *Rect) SetHeight(height int) {
	r.height = height
}

func (r *Rect) SetWidth(width int) {
	r.width = width
}

func (r *Rect) Height() int {
	return r.height
}

func (r *Rect) Width() int {
	return r.width
}

type Cube struct {
	Rect
	breadth int
}

func NewCube(side int) *Cube {
	r := Rect{}
	r.SetHeight(side)
	r.SetWidth(side)
	return &Cube{r, side}
}

func (c *Cube) SetBreadth(breadth int) {
	c.breadth = breadth
}

func (c *Cube) Breadth() int {
	return c.breadth
}

func (c *Cube) Volume() int {
	return c.Area() * c.Breadth()
}

func (r *Rect) Area() int {
	return r.Height() * r.Width()
}

func InterfaceSegregationPrinciple() {

	c := NewCube(10)
	fmt.Println("Volume", c.Volume())
}
