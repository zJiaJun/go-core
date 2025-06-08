package context

import (
	"fmt"
	"log"
	"sync"
)

func UseChanPipeline() {
	buy := pipelineBuy(100)
	build_1 := pipelineBuild(buy)
	build_2 := pipelineBuild(buy)
	build_3 := pipelineBuild(buy)
	merge := pipelineMerge(build_1, build_2, build_3)
	pack := pipelinePackage(merge)
	for p := range pack {
		log.Println(p)
	}
}

func pipelineBuy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprintf("[配件-%d]", i)
		}
	}()
	return out
}

func pipelineBuild(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			out <- fmt.Sprintf("[组装-%s]", s)
		}
	}()
	return out
}

func pipelinePackage(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			out <- fmt.Sprintf("[打包-%s]", s)
		}
	}()
	return out
}

func pipelineMerge(ins ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	p := func(in <-chan string) {
		defer wg.Done()
		for s := range in {
			out <- s
		}
	}
	for _, in := range ins {
		go p(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
