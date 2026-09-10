package main

import (
	"fmt"
	"sync"
	"time"
)

func Reader(jobCh chan string, resCh chan string, wg *sync.WaitGroup) {
	// create a workerpool pattern
	defer wg.Done()
	// after job done 1 gorutine

	for val := range jobCh {
		// range input channels in every gorutine
		(time.Sleep(100 * time.Millisecond))
		// imitation working
		resCh <- fmt.Sprintf("done: %s", val)
		// send in a resulting channel
	}

}

func main() {
	wg := sync.WaitGroup{}
	// create a WaitingGroup
	URL := []string{
		// create a slice URL adresses
		"http://page-hub.net/profile",
		"http://xzhbmwcpyri.net",
		"http://siteflow.com/main/052av",
		"https://spacex.dev/api/il7hj",
		"https://digitalsite.biz/index",
		"https://nexusnet.dev/api/chbtq",
		"https://ygxiskqmh.org",
		"https://ylwjfglafs.com",
		"http://wavex.cohttps://techflow.com/v1",
	}
	jobCh := make(chan string)
	// create input channel
	resCh := make(chan string)
	// create resulting channel
	wg.Add(3)
	go Reader(jobCh, resCh, &wg)
	go Reader(jobCh, resCh, &wg)
	go Reader(jobCh, resCh, &wg)
	// starting 3 bacground worker (workerpool)
	go func() {
		// background send in input channel our URLs
		for i := 0; i < len(URL); i++ {
			jobCh <- URL[i]
		}
		close(jobCh)
		// after close input channel
	}()
	go func() {
		// wait all workers
		wg.Wait()
		close(resCh)
		// & close the resulting channel
	}()
	for val := range resCh {
		// printing resulting
		fmt.Println(val)
	}
}
