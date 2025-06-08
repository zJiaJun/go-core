package context

import (
	"log"
	"time"
)

func UseChanFuture() {
	vegetablesCh := washVegetables()
	waterCh := boilWater()
	log.Println("已经安排洗菜和烧水了, 我先眯一会,做其他事情")
	time.Sleep(2 * time.Second)
	log.Println("开始做火锅了,看看菜和水好了吗")
	vegetables := <-vegetablesCh
	water := <-waterCh
	log.Printf("准备好了,可以做火锅了, %s, %s", vegetables, water)
}

func washVegetables() <-chan string {
	vegetables := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vegetables <- "洗好的菜"
	}()
	return vegetables
}

func boilWater() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		water <- "烧烤的水"
	}()
	return water
}
