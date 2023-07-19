package concurrency

import (
	"fmt"
	"time"
)

func MainChannel() {
    ch := make(chan string)
    go countCh(ch)
    go secondCh(ch)

    fmt.Println(<- ch)
    fmt.Println(<- ch)

    start := time.Now()

        elapsed := time.Since(start)
    fmt.Printf("Process it took %s seconds", elapsed)

}

func countCh(ch chan string) {
    fmt.Println("\n*** countCh INIT ***")
    for i := 0; i < 100; i++ {
        fmt.Printf("%d Mississipi...\n", i)
    }
    ch <- "countCh DONE\n"
}

func secondCh(ch chan string) {
    time.Sleep(time.Second*2)
    fmt.Println("*** secondCh INIT ***")
    ch <- "secondCh DONE\n"
}


