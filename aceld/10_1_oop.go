package main

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

// 声明一种新的数据类型myint,是int的别名
type myint int

// 定义一个结构体
type Book struct {
	title  string
	author string
}

func changeBook1(book Book) {
	//传递book的副本
	book.author = "asd"
}
func changeBook2(book *Book) {
	//指针传递
	book.author = "qwe"
}

func Oop101Example() {
	defer fmt.Println(common.Separator)
	var a myint = 10
	fmt.Println("a =", a)

	var book1 Book
	book1.title = "GoLang"
	book1.author = "leon"
	fmt.Printf("%v\n", book1)
	changeBook1(book1)
	fmt.Printf("%v\n", book1)
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}
