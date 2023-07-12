package concurrency

import (
	"fmt"
	"sync"
	"time"
)

// wait groups wait for a collection of concurrent tasks
func MainWaitGroup() {
    var wg sync.WaitGroup
    wg.Add(1)

    go func() {

        countWG("sheep")
        wg.Done()
    }()
    wg.Wait()


    fmt.Println("Second task")
} 

func countWG(item string) {
    for i := 1; i <= 5; i++ {
        fmt.Println(i, item)
        time.Sleep(time.Millisecond * 500)
    }

}
