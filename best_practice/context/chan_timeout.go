package context

import (
	"log"
	"sync"
	"time"
)

func UseChanTimeout() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		timeoutWork(ch)
	}()
	wg.Wait()
}

func timeoutWork(ch <-chan int) {
	select {
	case v := <-ch:
		log.Printf("chan receive, %d", v)
	case <-time.After(2 * time.Second):
		log.Println("timeout")
	}
}
