package main

import (
	"fmt"
)

type Human struct {
	sex string
	age int
}

type Labor struct {
	Human
	Job       string
	seniority int
}

func (labor Labor) String() string {
	return fmt.Sprintf("%s is a %d years old %s with %d years of experience", labor.Job, labor.age, labor.sex, labor.seniority)
}

func main() {
	m_labor := Labor{
		Human: Human{
			sex: "male",
			age: 20,
		},
		Job:       "programmer",
		seniority: 1,
	}

	fmt.Println(m_labor)
	fmt.Println(m_labor.String())

}
