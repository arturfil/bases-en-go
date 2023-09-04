package printer

type calculator interface {
    Average() float64
}

func Average(calc calculator) float64 {
    return calc.Average()
}


