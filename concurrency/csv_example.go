package concurrency

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type api struct {
    name string
   url string
}

var apis = []api {
    {"countries", "https://restcountries.com/v3.1/all"},
    {"pokemon", "https://pokeapi.co/api/v2/pokemon/?limit=151&offset=0"},
}

func CSVExampleMain() {

    go readCSV() // 1
    go readCSV() // 2 
    go readCSV() // 3
    go readCSV() // 4
    go readCSV() // 5
    go readCSV() // 5
    go readCSV() // 6
    go readCSV() // 7
    go readCSV() // 8
    go readCSV() // 9 
    go readCSV() // 10
    start := time.Now()

    elapsed := time.Since(start)
    fmt.Printf("Process it took %s seconds", elapsed)
}

func readApis() {

    for _, api := range apis {
    resp, err := http.Get(api.url)
    if err != nil {
        log.Fatal(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))
    }
    fmt.Println("") 

}

func readCSV() {
    f, _ := os.Open("emails_hondthou.csv")
    r := csv.NewReader(f)
    for {
        record, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }
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
