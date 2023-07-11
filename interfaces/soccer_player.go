package interfaces

import "fmt"

type SoccerPlayer struct {
    Name      string
    Team      string
    GameScore GameScore
    Goals     int
}

func NewSoccerPlayer(name string, team string) *SoccerPlayer {
    return &SoccerPlayer{
        Name:      name,
        Team:      team,
        GameScore: GameScore{Us: 0, Oposittion: 0},
        Goals:     0,
    }
}

func (s *SoccerPlayer) Score(points int) {
    s.Goals += points
    s.GameScore.Us += points
    fmt.Printf("%v just scored for %v\n", s.Name, s.Team)
    fmt.Printf("Score now is %v to %v\n", s.GameScore.Us, s.GameScore.Oposittion)
    fmt.Println("---")
}

func (s *SoccerPlayer) Defend(points int, success bool) {
    if !success {
        s.GameScore.Oposittion += points
        fmt.Printf("%v conceded a goal\n", s.Team)
        fmt.Printf("Score now is %v to %v\n", s.GameScore.Us, s.GameScore.Oposittion)
        fmt.Println("---")
    } else {
        fmt.Println("Successfull tackle")
        fmt.Println("---")
    }
}

