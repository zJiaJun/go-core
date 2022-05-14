package bilibili_aceld

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
	"time"
)

func Channel132Example() {
	defer fmt.Println(common.Separator)
	//定义一个channel, 带有缓冲
	//当channel容量满了, 再向里面写数据, 会阻塞
	//当channel为空, 从里面取数据, 也会阻塞
	c := make(chan int, 3)
	//len元素数量, cap 容量
	fmt.Println("len(c) = ", len(c), ",cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("child goroutine finished")
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("child goroutine running, send = ", i,
				",len(c) = ", len(c),
				",cap(c)", cap(c))
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("main goroutine num = ", num)
	}
	fmt.Println("main goroutine finished")
}
