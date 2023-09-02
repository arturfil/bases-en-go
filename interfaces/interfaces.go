package interfaces

import "fmt"

/*
    The Group struct and the Player structs are two different structs for an application that
    tracks data for a soccer related application

    Both structs have fields where we could get averages
    As such, I created an interface that shares the expected action Average  

    With this I can create a "generic" function or a shared behavior/method between these two structs
*/

type Group struct {
    grades[] int
    name string
    admin string
}

type Player struct {
    reviews[] int
    name string
    position string
}

type Average interface {
   Average() float32 
}

func (g *Group) Average() float32 {
    var sum float32 = 0
    for _, grade := range g.grades {
        sum += float32(grade) 
    }
    return sum / float32(len(g.grades)) // average computed
}

func (p *Player) Describe() {
    fmt.Println(
        "This player is", p.name, 
        "and plays in position", p.position,
    )
}

func (p *Player) Average() float32 {
    var sum float32 = 0
    for _, val := range p.reviews {
        sum += float32(val)
    }
    return sum / float32(len(p.reviews))
}

func PrintAverage(a Average) {
    fmt.Printf("The avg is: %.2f\n", a.Average())
}

func InterfacesMain() {

    rooftopGroup := &Group{
        grades: []int {2,3,4,5,6,7,7,8,9},
        name: "Rooft Top Soccer Group",
        admin: "Chamo",
    } 

    arturo := &Player{
        reviews: []int {1,2,3,6,7,8,6,5,4,4},
        name: "Arturo", 
        position: "attacking midfield",
    }

    arturo.Describe() 

    PrintAverage(arturo)
    PrintAverage(rooftopGroup)

}
