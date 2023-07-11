package loops

import (
	"fmt"

	arraysslices "github.com/artufil/yt-bases/arrays_slices"
)

var games = arraysslices.Videogames

func LoopsMain() {
    // for _, game := range(games) {
    //     fmt.Println(game)
    // } 

    // for i := 0; i < len(games); i++ {
    //     fmt.Println(games[i]) 
    // }

    // while loop
    k := 0 
    for k < len(games) {
        fmt.Println(games[k])
        k++
    }
    
}
