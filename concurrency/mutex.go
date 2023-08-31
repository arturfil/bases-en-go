package concurrency

import (
	"fmt"
	"sync"
)

var (
    mutex sync.Mutex
    balance int = 0

    retiros = []int {
        500,
        200,
    }

    depositos = []int {
        500,
        200,
    }

)

func MutexMain() {
    var wg sync.WaitGroup

    for _, trx := range depositos {
        wg.Add(1)
        go func(monto int) {
            defer wg.Done()
            for i := 0; i < 1000; i++ {
                deposito(monto)
            }
        }(trx)
    } 

    for _, trx := range retiros {
        wg.Add(1)
        go func(monto int) {
            defer wg.Done()
            for i := 0; i < 1000; i++ {
                retiro(monto)
            }
        }(trx)
    } 

    wg.Wait()

    fmt.Println("----------------------------------------------------------")
    fmt.Printf("New Balance %d\n", balance)

}

func deposito(amount int) {
    mutex.Lock()
    defer mutex.Unlock()

    balance += amount
    fmt.Printf("[Deposito]\t | %d\t | to account\t | with balance %d\n", amount, balance)

}

func retiro(amount int) {
    mutex.Lock()
    defer mutex.Unlock()

    balance -= amount
    fmt.Printf("[Retiro]\t | %d\t | to account\t | with balance %d\n", amount, balance)
}
