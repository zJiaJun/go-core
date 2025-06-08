package context

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

func UseContextStopMultipleGoroutine() {
	var wg sync.WaitGroup
	wg.Add(5)
	ctx, cancel := context.WithCancelCause(context.Background())
	go func() {
		defer wg.Done()
		go func() {
			defer wg.Done()
			multiWatch(ctx, "dog_1_1")
		}()
		multiWatch(ctx, "dog_1")
	}()
	go func() {
		defer wg.Done()
		go func() {
			defer wg.Done()
			multiWatch(ctx, "dog_2_1")
		}()
		multiWatch(ctx, "dog_2")

	}()
	go func() {
		defer wg.Done()
		multiWatch(ctx, "dog_3")
	}()
	time.Sleep(5 * time.Second)
	cancel(errors.New("cancel"))
	wg.Wait()
}

func multiWatch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("%s is watching stop, error: %v, cause: %v", name, ctx.Err(), context.Cause(ctx))
			return
		default:
			log.Printf("%s is watching", name)
		}
		time.Sleep(1 * time.Second)
	}
}
