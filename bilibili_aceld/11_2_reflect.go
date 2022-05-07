package bilibili_aceld

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
	"reflect"
)

func reflectNum(arg interface{}) {
	fmt.Println("type:", reflect.TypeOf(arg))
	fmt.Println("value:", reflect.ValueOf(arg))
}

type reflectUser struct {
	id   int
	name string
	age  int
}

func (this reflectUser) Call() {
	fmt.Println("reflect user call")
	fmt.Printf("%v\n", this)
}

func reflectUserVal(user interface{}) {
	ut := reflect.TypeOf(user)
	fmt.Println("user type:", ut)
	uv := reflect.ValueOf(user)
	fmt.Println("user value:", uv)

	for i := 0; i < ut.NumField(); i++ {
		field := ut.Field(i)
		value := uv.Field(i)
		fmt.Println("field = ", field.Name, ",value = ", value)
	}

	for i := 0; i < ut.NumMethod(); i++ {
		m := ut.Method(i)
		fmt.Println("method name = ", m.Name, ",type = ", m.Type)
	}

}

func Reflect112Example() {
	defer fmt.Println(common.Separator)
	var num float64 = 1.2345
	reflectNum(num)

	ru := reflectUser{1, "li4", 20}
	ru.Call()
	reflectUserVal(ru)
}
