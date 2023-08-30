package concurrency

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func WaitGroupsMain() {
    var wg sync.WaitGroup

	start := time.Now()

    /*
    wg.Add(3) 

    go loopNumbers(&wg, 0, 100)	
    go loopNumbers(&wg, 100, 200)	
    go loopNumbers(&wg, 200, 300)	

    wg.Wait()
    */

    urls := []string {
        "http://google.com",
        "http://youtube.com",
        "http://twitter.com",
        "http://facebook.com",
    }

    getStatus(&wg, urls)

	elapsed := time.Since(start)

    fmt.Println("Representing another process")
    
	fmt.Printf("Time elapsed %fs\n", float32(elapsed)/1000000000)

}

func loopNumbers(wg *sync.WaitGroup, init int, end int) {
    time.Sleep(time.Second * 3)
    for i := init; i < end; i++ {
        fmt.Println(i)
    }
    wg.Done()
}

func getStatus(wg *sync.WaitGroup, urls []string) {
    for _, url := range urls {
        wg.Add(1)
        go func(url string) {
            res, err := http.Get(url)

            if err != nil {
                fmt.Printf("[Error: %s x -> %s]\n", err, url) 
            } else {
                fmt.Printf("[%d %s]\n", res.StatusCode, url)    
            }
            wg.Done()
        }(url)
    }
    wg.Wait()
    fmt.Println("All goroutines are done!")
}
