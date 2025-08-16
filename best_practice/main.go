package main

import (
	"github.com/zjiajun/go-core/best_practice/reflect"
	"time"
)

func main() {
	//context.UseContextStopGoroutine()
	//context.UseChanStopGoroutine()
	//context.UseContextStopMultipleGoroutine()
	//context.UseChanTimeout()
	//context.UseChanPipeline()
	//context.UseChanFuture()
	//reflect.SetVal()
	//reflect.SetStructFieldValue()
	reflect.GetKind()
	time.Sleep(1 * time.Hour)
}
