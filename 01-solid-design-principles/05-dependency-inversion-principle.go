package main

import "fmt"

/*************************************************************************************
Dependency Inversion Principle
The Dependency Inversion Principle (DIP) states that high level modules should not depend on low level modules;
 both should depend on abstractions. Abstractions should not depend on details.
 Details should depend upon abstractions. It’s extremely common when writing software to implement it such that
 each module or method specifically refers to its collaborators, which does the same.
 This type of programming typically lacks sufficient layers of abstraction, and results in a very tightly coupled system,
 since every module is directly referencing lower level modules.
*/

type Author struct {
	name string
	// other useful stuff here
}

type Journal struct {
	title  string
	author Author
}

// JournalBrowser forms an abstract layer
type JournalBrowser interface {
	FindAllJournalsOf(name string) []*Journal
}

// Journals is a low-level component as it holds some data
// This data can come fron a database, file etc.
type Journals struct {
	journals []Journal
}

// AddJournal adds a new journal entry
func (j *Journals) AddJournal(journal Journal) {
	j.journals = append(j.journals, journal)
}

// FindAllJournalsOf forms an abstract layer to filter data from all journals
// The caller component need not to be aware of the filteration mechanism.
// So, if this method is modified, no change is required in the caller function / component.
func (j *Journals) FindAllJournalsOf(name string) []*Journal {
	result := make([]*Journal, 0)

	for i, v := range j.journals {
		if v.author.name == name {
			result = append(result, &j.journals[i])
		}
	}

	return result
}

// High - level component as it processess data
type Research struct {

	// journals Journals

	// Rather than accessing Journals directly (As shown in above statement)
	// Use an abstract layer JournalBrowser to access data
	browser JournalBrowser // low-level
}

func (r *Research) Investigate(name string) {
	// usage of FindAllJournalsOf facilitates loose coupling
	// Therefore, if the mechanism of filtering record changes in the called function,
	// there is no need to modify this function.
	for _, j := range r.browser.FindAllJournalsOf(name) {
		fmt.Printf("%s is author of %s\n", name, j.title)
	}
}

func DependencyInversionPrinciple() {

	journals := Journals{}
	journals.AddJournal(Journal{"Journal 1", Author{"John"}})
	journals.AddJournal(Journal{"Journal 2", Author{"John"}})
	journals.AddJournal(Journal{"Journal 3", Author{"Jane"}})
	journals.AddJournal(Journal{"Journal 4", Author{"Marry"}})
	journals.AddJournal(Journal{"Journal 5", Author{"Peter"}})

	research := Research{&journals}
	research.Investigate("John")
}
