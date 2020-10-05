package main

type Teacher interface {
	Name() string
	Qualification() string
	SetName(string)
	SetQualification(string)
}

type teacher struct {
	name          string
	qualification string
}

func (t *teacher) Name() string {
	return t.name
}

func (t *teacher) Qualification() string {
	return t.qualification
}

func (t *teacher) SetName(name string) {
	t.name = name
}

func (t *teacher) SetQualification(qualification string) {
	t.qualification = qualification
}

// Interface factory example
func NewTeacher(name string, qualification string) Teacher {
	return &teacher{name, qualification}
}
