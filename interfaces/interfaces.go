package interfaces

import (
	"fmt"

	"github.com/artufil/yt-bases/groups"
	"github.com/artufil/yt-bases/players"
	"github.com/artufil/yt-bases/printer"
)

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

    rooftopGroup := groups.New([]int {2,2,3,2,2}, "Rooftop", "Juan")
    arturo := players.New([]int {3,2,3,3,3,2}, "Arturo", "Midfield")

    fmt.Printf("%.2f\n", printer.Average(rooftopGroup))
    fmt.Printf("%.2f\n", printer.Average(arturo))

}
