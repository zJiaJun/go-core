package aceld

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

type human struct {
	name string
	sex  string
}

func (h *human) Eat() {
	fmt.Println("Human.Eat()...")
}
func (h *human) Walk() {
	fmt.Println("Human.Walk()...")
}

type superMan struct {
	human //superMan类继承了human类的字段和方法
	level int
}

//重定义父类的方法Eat()
func (s *superMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

//子类添加新方法
func (s *superMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func Oop103Example() {
	defer fmt.Println(common.Separator)
	h := human{name: "zhang3", sex: "female"}
	h.Eat()
	h.Walk()

	sm := superMan{human: human{"li4", "female"}, level: 10}
	sm.Eat()  //子类的方法
	sm.Fly()  //子类的方法
	sm.Walk() //父类的方法

}
