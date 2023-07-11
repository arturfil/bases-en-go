package concurrency

import (
	"fmt"
	"time"
)

func GoRoutinesMain() {
    go count("sheep")    
    go count("fish")

    time.Sleep(time.Second * 2)
}

func count(item string) {
    for i := 1; true; i++ {
        fmt.Println(i, item)
        time.Sleep(time.Millisecond * 500)
    }
}
