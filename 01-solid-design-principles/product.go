package main

const (
	// Colors list
	_ Color = iota
	Red
	Blue
	Green
	Brown
)

const (
	// Size list
	_ Size = iota
	Large
	Medium
	Small
)

// Color of the product
type Color int

// Size of the product
type Size int

// Product hold the data of a product
type Product struct {
	ProductTitle string
	ProductSize  Size
	ProductColor Color
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) SetTitle(title string) *Product {
	p.ProductTitle = title
	return p
}

func (p *Product) SetSize(size Size) *Product {
	p.ProductSize = size
	return p
}

func (p *Product) SetColor(color Color) *Product {
	p.ProductColor = color
	return p
}

func (p *Product) Title() string {
	return p.ProductTitle
}

func (p *Product) Size() Size {
	return p.ProductSize
}

func (p *Product) Color() Color {
	return p.ProductColor
}
