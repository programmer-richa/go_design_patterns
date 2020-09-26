package main

import (
	"fmt"
)

/***********************************************************
* Separate of concerns
 According to this principle, a component is responsible
  In the following example Book struct is responsible for holding information of a book
  A separate JsonHelper struct is responsible to write Json Data.
  Therefore, instead of attaching ToJSON method with book, a separate component is created for that functionality.
*/

func singleResponsibilityPrinciple() {
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
