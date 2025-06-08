package context

import (
	"log"
	"sync"
	"time"
)

func UseChanStopGoroutine() {
	var wg sync.WaitGroup
	ch := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		chanWatch(ch, "dog")
	}()
	time.Sleep(5 * time.Second)
	ch <- true
	wg.Wait()
}

func chanWatch(ch <-chan bool, name string) {
	for {
		select {
		case <-ch:
			log.Printf("%s is watching stop", name)
			return
		default:
			log.Printf("%s is watching", name)
		}
		time.Sleep(1 * time.Second)
	}
}
