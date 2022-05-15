package bilibili_aceld

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

func fib(c, quit chan int) {
	x, y := 1, 1
	for {
		//select 具备多路channel的监控状态功能
		select {
		case c <- x:
			x = y
			y = x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Channel134Example() {
	defer fmt.Println(common.Separator)
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fib(c, quit)
}
