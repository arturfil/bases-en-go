package interfaces

// Notes are private

type TeamPlayer interface {
   Score(points int) 
   Defend(points int, success bool)
}

type GameScore struct {
    Us int
    Oposittion int
}

type Rectangle struct {
    height, base int
}

type Square struct {
    edge int 
}

func (s *Square) calculateArea() float32 {
    area := float32(s.edge * s.edge)
    return area
}

func (r *Rectangle) calculateArea() float32 {
    area := float32(r.base * r.height)
    return area
}

func InterfaceMain() {
    var cristiano TeamPlayer
    cristiano = NewSoccerPlayer("Cristiano Ronaldo", "Manchester United")
    cristiano.Score(1)
    cristiano.Defend(1, false)

    var jimmyButter TeamPlayer = NewBasketBallPlayer("Jimmy Buttler", "Miami Heat")
    jimmyButter.Score(3)
    jimmyButter.Defend(2, true)

}

