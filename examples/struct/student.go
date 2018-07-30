package main

import (
	"fmt"
)

type Mark struct {
	english int64
	maths   int64
	science int64
	social  int64
}

type Student struct {
	id   string
	name string
	Mark
}

func (s *Student) setName(name string) {
	s.name = name
}

func (s *Student) setId(id string) {
	s.id = id
}

func (s *Student) setMark(marks Mark) {
	s.Mark = marks
}

func (s *Student) getStudentMark() {
	fmt.Printf("English :: %d\n", s.english)
	fmt.Printf("Maths :: %d\n", s.maths)
	fmt.Printf("Science :: %d\n", s.science)
	fmt.Printf("Social :: %d\n", s.social)
}

func main() {
	m := Mark{
		english: 90,
		maths:   80,
		science: 70,
		social:  65,
	}

	s := &Student{
		id:   "#1242",
		name: "phani",
		Mark: m,
	}

	m1 := m
	m1.english = 85

	s.setMark(m1) // mutate the Mark value
	s.getStudentMark()

}
