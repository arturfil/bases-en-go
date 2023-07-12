package concurrency

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func MainChannel() {
    /*
    ch := make(chan string)
    start := time.Now()
    go countCh(ch)
    go secondCh(ch)

    fmt.Println(<- ch)
    fmt.Println(<- ch)
    elapsed := time.Since(start)
    fmt.Printf("Process it took %s seconds", elapsed)
    */
    readCSV()
}

/*func countCh(ch chan string) {
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
}*/

func readCSV() {
    f, _ := os.Open("myFile0.csv")
    r := csv.NewReader(f)
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }
        // fmt.Print("[")
        for _, val := range record {
            if len(val) > 18 {
                fmt.Printf(" | %-28v", val)
            } else if len(val) == 1 {
                fmt.Printf(" | %-4v", val)
            } else {
                fmt.Printf(" | %-10v", val)
            }
        }
        fmt.Printf("|\n")

    }
    
}
