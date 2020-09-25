package main

import "strings"

// Book struct is responsible for holding data regarding a book
type Book struct {
	BookTitle   string   `json:"title"`
	BookAuthors []string `json:"authors"`
	PublishYear int      `json:"year"`
}

func NewBook() *Book {
	return &Book{}
}

func (b *Book) SetTitle(title string) *Book {
	b.BookTitle = title
	return b
}

func (b *Book) AddAuthor(author string) *Book {
	b.BookAuthors = append(b.BookAuthors, author)
	return b
}

func (b *Book) SetYear(year int) *Book {
	b.PublishYear = year
	return b
}

func (b *Book) Title() string {
	return b.BookTitle
}

func (b *Book) Author() string {
	return strings.Join(b.BookAuthors, ",")
}

func (b *Book) Year() int {
	return b.PublishYear
}
