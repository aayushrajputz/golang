package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func(){
		wg.Done()
		p.mu.Unlock()
	}()


	p.mu.Lock()
	p.views += 1

	

}

func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)

	}

	wg.Wait()
	fmt.Println(myPost.views)
}
