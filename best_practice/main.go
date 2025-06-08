package main

import (
	"github.com/zjiajun/go-core/best_practice/context"
	"time"
)

func main() {
	//context.UseContextStopGoroutine()
	//context.UseChanStopGoroutine()
	//context.UseContextStopMultipleGoroutine()
	//context.UseChanTimeout()
	//context.UseChanPipeline()
	context.UseChanFuture()
	time.Sleep(1 * time.Hour)
}
