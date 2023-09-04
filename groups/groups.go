package groups

type PlayerSet struct {
    grades []int
    name string
    admin string
}

func New(grades []int, name, admin string) *PlayerSet {
    return &PlayerSet{
        grades: grades, 
        name: name,
        admin: admin,
    }
}

func (ps *PlayerSet) Average() float64 {
    s := 0

    for _, g := range ps.grades {
        s += g
    }
    return float64(s) / float64(len(ps.grades))
}

