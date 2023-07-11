package pointers

import "fmt"

type soccerPlayer struct {
    name string
    team string
}

func initSoccerPlayer(name string, team string) *soccerPlayer {
    return &soccerPlayer{name: name, team: team}
}


func PointersMain() {
    name := "Arturo" 
    messi := initSoccerPlayer("Lionel Messi", "PSG")
    fmt.Println(*messi)

    // fmt.Println(&name) // & address -> memory address
    
    // declaracion de variables
    name_address := &name // pointer #1
    fmt.Println(&name)

    var name_ptr *string = &name // usando el * estamos haciendo un dereference 
    var name_ptr2 *string = &name

    fmt.Println("---------------------------------------------------------")
    fmt.Printf("name_address\t%-15p \tvar: name_address\n", name_address)
    fmt.Printf("name_ptr\t%-15p \tvar: name_ptr\n", name_ptr) // dereference -> 
    fmt.Printf("name_ptr2\t%-15p \tvar: name_ptr2\n", name_ptr2)

    fmt.Println("---------------------------------------------------------")
    *name_ptr = "Santiago"
    fmt.Printf("My name is\t%-15s \tvar: name\n", name)

    *name_ptr2 = "Armin Van Bureen"
    fmt.Printf("My name is\t%-15s \tvar: name\n", name)

    *name_address = "Jorge"
    fmt.Printf("My name is\t%-15s \tvar: name\n", name)
    
}
