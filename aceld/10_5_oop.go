package main

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

// interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc called", arg)
	// interface{} 如何区分, 到底是什么类型?

	value, ok := arg.(string)
	if ok {
		fmt.Println("arg is string type, value =", value)
	} else {
		fmt.Println("arg is not string type")
	}

	switch val := arg.(type) {
	case book:
		fmt.Println("myFunc book ", val)
	case int:
		fmt.Println("myFunc int", val)
	case string:
		fmt.Println("myFunc string", val)
	case float64:
		fmt.Println("myFunc float64", val)
	case bool:
		fmt.Println("myFunc bool", val)
	default:
		fmt.Println("myFunc type error", val)
	}
}

type book struct {
	auth string
}

func Oop105Example() {
	defer fmt.Println(common.Separator)
	myFunc(book{"zhang3"})
	myFunc(100)
	myFunc("str")
	myFunc(3.14)
	myFunc(true)
	myFunc('x')
}
