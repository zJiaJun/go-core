package goweb

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

type person struct {
	human
	Name  string
	Age   int
	Phone int
}

type human struct {
	Weight int
	Phone  int
}

func init() {
	fmt.Println("struct init running...")
}

func createPerson() *person {
	var P person
	P.human = human{100, 13012345678}
	P.Name = "Jack"
	P.Age = 10
	P.Phone = 123

	Q := person{human{110, 13212345678}, "Leon", 10, 456}
	//Q.Phone = 456
	//Q.human.Phone = 13012345678
	Q.Name, Q.Age = "Leon", 20

	W := person{
		human: human{Weight: 120, Phone: 13912345678},
		Name:  "Tom", Age: 30, Phone: 789}

	E := new(person)
	fmt.Println("StructExample", P)
	fmt.Println("StructExample", Q)
	fmt.Println("StructExample", W)
	fmt.Println("StructExample", E)
	return E
}

func setPersonValue(person *person, name string, age int) {
	person.human = human{130, 18812345678}
	person.Name = name
	person.Age = age
	person.Phone = 912
	fmt.Println("StructExample SetPersonValue person", person)
}
func (h *human) sayHi() {
	fmt.Println("StructExample,human sai hi", h.Weight, h.Phone)
}

//method重写,person重写human的sayHi
func (p *person) sayHi() {
	fmt.Println("StructExample,person sai hi", p.Name, p.human)
}

func structExample() {
	defer fmt.Println(common.Separator)
	person := createPerson()
	setPersonValue(person, "newName", 123)
	person.sayHi()
}
