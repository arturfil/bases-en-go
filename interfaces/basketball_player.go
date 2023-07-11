package interfaces

import "fmt"

type BasketBallPlayer struct {
    Name     string
    Points   int
    Rebounds int
    GameScore GameScore
    Team string
}

func NewBasketBallPlayer(name string, team string) *BasketBallPlayer {
    return &BasketBallPlayer{
        Name: name,
        Team: team,
        Rebounds: 0,
        GameScore: GameScore{Us: 0, Oposittion: 0},
        Points: 0,
    }
}

func (b *BasketBallPlayer) Score(points int) {
    b.Points += points
    b.GameScore.Us += points
    fmt.Printf("%v just scored for %v\n", b.Name, b.Team)
    fmt.Printf("Score now is %v to %v\n", b.GameScore.Us, b.GameScore.Oposittion)
    fmt.Println("---")
}

func (b *BasketBallPlayer) Defend(points int, success bool) {
    if !success {
        b.GameScore.Oposittion += points
        fmt.Println("Block was not successfull")
        fmt.Printf("Game Score is %v to %v\n", b.GameScore.Us, b.GameScore.Oposittion)
    } else {
        b.Rebounds += 1
        fmt.Println("Successfully blocked the shot attempt")
    }
}

