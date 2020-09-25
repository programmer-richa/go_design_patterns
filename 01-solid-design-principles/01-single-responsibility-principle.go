package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

/***********************************************************
* According to this principle, a component is responsible
  In the following example Book struct is responsible for holding information of a book
  A separate JsonHelper struct is responsible to write Json Data.
  Therefore, instead of attaching ToJSON method with book, a separate component is created for that functionality.
*/

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

// JsonHelper is responsible for json manipulation operations
type JsonHelper struct {
	data interface{}
}

// ToJSON converts data to JSON format
// data is the struct / map / slice type variable
func (j *JsonHelper) ToJSON() (s []byte, err error) {
	js, err := json.Marshal(j.data)
	if err != nil {
		return nil, err
	}
	var dat interface{}
	if err = json.Unmarshal(js, &dat); err != nil {
		return nil, err
	}
	mapB, err := json.Marshal(dat)
	if err != nil {
		return nil, err
	}
	return mapB, nil
}
func main() {
	// initializing 2 variables of type book
	b1 := NewBook().
		SetTitle("Book 1").
		AddAuthor("Author 1").
		AddAuthor("Author 2").
		SetYear(2000)

	b2 := NewBook().
		SetTitle("Book 2").
		AddAuthor("Author 1").
		AddAuthor("Author 2").
		SetYear(2010)

	// Initializing a slice of books
	books := []*Book{b1, b2}
	// Coverting slice of books to JSON
	jh := JsonHelper{data: books}
	b, err := jh.ToJSON()
	if err != nil {
		// print error in case some error occurs
		fmt.Println("Error occured while creating json: ", err.Error())
	} else {
		// print json string of slice of books
		fmt.Println(string(b))
	}

}
