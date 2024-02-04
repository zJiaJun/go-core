package main

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
	"time"
)

func newTask() {
	for i := 0; i < 3; i++ {
		fmt.Println("new goroutine, i =", i)
		time.Sleep(1 * time.Second)
	}
}

/*
B
B.defer
A
A.defer
*/
func anonymousFunc() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			//退出当前的goroutine
			//runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
}

func Goroutine12Example() {
	defer fmt.Println(common.Separator)
	//创建一个go协程
	go newTask()
	for i := 0; i < 3; i++ {
		fmt.Println("main goroutine, i =", i)
		time.Sleep(1 * time.Second)
	}
	anonymousFunc()

}
