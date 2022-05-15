package aceld

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

func deferFun1() {
	fmt.Println("1")
}
func deferFun2() {
	fmt.Println("2")
}
func deferFun3() {
	fmt.Println("3")
}

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}
func returnFunc() int {
	fmt.Println("return func called")
	return 0
}
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

/*
	1.defer 执行顺序, 先进后出
	deferFun1 入栈
	deferFun2 入栈
	deferFun3 入栈

	deferFun3 出栈
	deferFun2 出栈
	deferFun1 出栈

	2.先return执行, 后defer执行
*/
func DeferExample() {
	defer fmt.Println(common.Separator)
	defer deferFun1()
	defer deferFun2()
	defer deferFun3()
	returnAndDefer()
	fmt.Println("deferExample running...")
}
