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
func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "Cat"
}

//具体的类
type Dog struct {
	color string
}

//实现了AnimalInterface的三个方法
func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}
func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) GetType() string {
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
