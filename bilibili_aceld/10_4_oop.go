package bilibili_aceld

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

//接口,本质是一个指针
type AnimalInterface interface {
	Sleep()
	GetColor() string
	GetType() string
}

//具体的类
type Cat struct {
	color string
}

//实现了AnimalInterface的三个方法
func (c *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}
func (c *Cat) GetColor() string {
	return c.color
}
func (c *Cat) GetType() string {
	return "Cat"
}

//具体的类
type Dog struct {
	color string
}

//实现了AnimalInterface的三个方法
func (d *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}
func (d *Dog) GetColor() string {
	return d.color
}
func (d *Dog) GetType() string {
	return "Dog"
}
func showAnimal(animal AnimalInterface) {
	//多态
	fmt.Println(animal.GetType())
}
func Oop104Example() {
	defer fmt.Println(common.Separator)
	var animal AnimalInterface //接口数据类型,父类指针
	//接口本质是一个指针, 所以需要& 传递地址
	animal = &Cat{"green"}
	animal.Sleep() //调用Cat的sleep方法
	animal = &Dog{"yellow"}
	animal.Sleep() //调用Dog的sleep方法

	cat := Cat{"green"}
	dog := Dog{"yellow"}
	showAnimal(&cat)
	showAnimal(&dog)
}
