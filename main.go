package main

import (
	"fmt"
	"github.com/zjiajun/go-core/aceld"
	"github.com/zjiajun/go-core/common"
	"github.com/zjiajun/go-core/goweb"
)

func main() {
	fmt.Println("[go web example begin]")
	goweb.Run()
	fmt.Println(common.Separator)
	fmt.Println("[bilibili aceld example begin]")
	aceld.Run()
}
