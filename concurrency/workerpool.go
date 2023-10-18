package concurrency

import (
	"fmt"
	"time"

)

func WorkersMain() {
   const numJobs = 10 
   start := time.Now()

   completed := []int{}

   trabajos := make(chan int, numJobs)
   trabajosCmpl := make(chan int, numJobs)

   for i := 1; i <= 3; i++ {
       go worker(i, trabajos, trabajosCmpl)
   }

   // blocking code, cargar los trabajos por hacer del 1-10; numJobs = 10
   for j := 1; j <= numJobs; j++ {
       trabajos <- j
   }

   close(trabajos)

   for k := 1; k <= numJobs; k++ {
       val := <- trabajosCmpl
       completed = append(completed, val) 
   }

   end := time.Since(start)
   fmt.Println("completedJobsChan", completed)
   fmt.Println("It has taken: ", end)

}
                            // read only            // write only
func worker(id int, trabajos <- chan int, trabajosCmpl chan <- int) {
    for job := range trabajos {
        time.Sleep(time.Second * 1)
        fmt.Println("worker\t", id, "\t[started]\t job", job, "with", len(trabajos), "jobs left to process")
        fmt.Println("worker\t", id, "\t[finished]\t job", job)
        trabajosCmpl <- job * job
    }
}
