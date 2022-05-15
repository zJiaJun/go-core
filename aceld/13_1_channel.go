package aceld

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

func Channel131Example() {
	defer fmt.Println(common.Separator)
	//定义一个channel, 无缓冲
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine finished")
		fmt.Println("goroutine running")
		//将666发送到channel中
		c <- 666
	}()

	//从channel中接收数据,并赋值给num
	num := <-c
	fmt.Println("num = ", num)
}
