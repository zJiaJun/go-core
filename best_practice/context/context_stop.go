package context

import (
	"context"
	"log"
	"sync"
	"time"
)

func UseContextStopGoroutine() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		contextWatch(ctx, "dog")
	}()
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}

func contextWatch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("%s is watching stop", name)
			return
		default:
			log.Printf("%s is watching", name)
		}
		time.Sleep(1 * time.Second)
	}
}
