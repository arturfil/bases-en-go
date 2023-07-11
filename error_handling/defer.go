package errorhandling

import (
	"fmt"
	"log"
	"net/http"
)

// defer, panic, recover
func showDefer() {
    fmt.Println("inicio")
    defer fmt.Println("procesando")
    fmt.Println("fin")
}

func showPanic() {
    fmt.Println("inicio")
    panic("Error, algo no esta funcionando")
    fmt.Println("fin")
}

func showRecover() {
    fmt.Println("inicio")
    defer func() {
        if err := recover(); err != nil {
            log.Println("Error", err)
        }
    }()
    panic("Algo salio mal")
    fmt.Println("Terminando funcion")
}

func ErrorMain() {
    ListenServer()
}

func ListenServer() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Aprendiendo Go!"))
    })
    fmt.Println("Listening on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err.Error())
    }
}
