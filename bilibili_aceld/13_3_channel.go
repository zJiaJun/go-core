package bilibili_aceld

import (
	"fmt"
	"github.com/zjiajun/go-core/common"
)

func Channel133Example() {
	defer fmt.Println(common.Separator)
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//close关键字, 关闭channel
		close(c)
	}()

	/*
		for {
			//ok如果为true表示channel没有关闭, false表示关闭
			if data, ok := <-c; ok {
				fmt.Println("data = ", data)
			} else {
				break
			}
		}
	*/
	//可以使用range来迭代不断操作channel
	for data := range c {
		fmt.Println("data = ", data)
	}
	fmt.Println("main goroutine finished")
}
