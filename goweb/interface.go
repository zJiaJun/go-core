package goweb

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
	"strconv"
)

type humans struct {
	name  string
	age   int
	phone string
}
type student struct {
	humans //匿名字段
	school string
	loan   float32
}

type employee struct {
	humans  //匿名字段
	company string
	money   float32
}

// humans 实现 sayHi方法
func (h humans) sayHi() {
	fmt.Println("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// humans 实现 sing 方法
func (h humans) sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// employee 重载 humans 的 sayHi方法
func (e employee) sayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

type men interface {
	sayHi()
	sing(lyrics string)
}

type element interface{}

type list []element
type per struct {
	name string
	age  int
}

func (p per) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}
func interfaceVarExample() {
	list := make(list, 3)
	list[0] = 1
	list[1] = "hello"
	list[2] = per{"leon", 30}

	for i := range list {
		value, ok := list[i].(int)
		fmt.Println(value)
		fmt.Println(ok)
	}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case per:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}

func interfaceExample() {
	defer fmt.Println(common.Separator)
	mike := student{humans{"Mike", 25, "222-222-XXX"},
		"MIT", 0.00}
	paul := student{humans{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := employee{humans{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := employee{humans{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	var i men
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.sayHi()
	i.sing("November rain")

	i = tom
	fmt.Println("This is tom, an Employee:")
	i.sayHi()
	i.sing("Born to be wild")

	// 定义了 slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]men, 3)
	// 这三个都是不同类型的元素，但是他们实现了 interface 同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.sayHi()
	}
}
