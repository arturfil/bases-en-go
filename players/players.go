package players

type Player struct {
	reviews  []int
	name     string
	position string
}

func New(reviews []int, name string, position string) *Player {
    return &Player{reviews, name, position}
}

func (p *Player) Average() float64 {
    s := 0
    for _, review := range p.reviews {
        s += review
    }
    return float64(s) / float64(len(p.reviews))
}

